package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"k8s.io/klog/v2"
	"strconv"
)

// SinkConfig represents the configuration for the Mongo sink connector.
// zh: SinkConfig 表示 Mongo 接收器连接器的配置。
type SinkConfig struct {
	CollectionName            string
	CollectionCapMaxDocuments int64
	CollectionCapMaxSizeBytes int64
	CollectionCapEnable       bool
}

// MongoSink represents an Mongo sink connector.
type MongoSink struct {
	ctx  context.Context
	conf SinkConfig
	db   *mongo.Database
	in   chan any
}

// NewMongoSink returns a new MongoSink instance.
// zh: NewMongoSink 返回一个新的 MongoSink 实例。
func NewMongoSink(ctx context.Context, db *mongo.Database, conf SinkConfig) (*MongoSink, error) {
	sink := &MongoSink{
		ctx:  ctx,
		conf: conf,
		db:   db,
		in:   make(chan any),
	}

	go sink.init()
	return sink, nil
}

// init starts the main loop.
func (ms *MongoSink) init() {
	ms.capCollection()
	for msg := range ms.in {
		// 处理msg
		//message := msg.(kafka.Message)

		// Insert the message into the collection.
		_, err := ms.db.Collection(ms.conf.CollectionName).InsertOne(context.Background(), msg)
		if err != nil {
			klog.ErrorS(err, "Problem inserting to mongo collection")
		}

		// 测试查询的数据
		//id, _ := primitive.ObjectIDFromHex("66d6d0aedd6a0497f2bfef31")
		//// 查询数据
		//res := ms.db.Collection(ms.conf.CollectionName).FindOne(context.Background(), bson.M{"_id": id})
		//if res.Err() != nil {
		//	klog.ErrorS(err, "Problem inserting to mongo collection")
		//}
		//var message kafka.Message
		//err = res.Decode(&message)
		//if err != nil {
		//	klog.ErrorS(err, "Problem inserting to mongo collection")
		//}
		//fmt.Println("message received: ", string(message.Value))
		//klog.InfoS("FindOne", "res", res)
	}
}

func (ms *MongoSink) capCollection() (ok bool) {
	colName := ms.conf.CollectionName
	colCapMaxSizeBytes := ms.conf.CollectionCapMaxSizeBytes
	colCapMaxDocuments := ms.conf.CollectionCapMaxDocuments
	colCapEnable := ms.conf.CollectionCapEnable

	if !colCapEnable {
		return false
	}

	exists, err := ms.collectionExists(colName)
	if err != nil {
		klog.ErrorS(err, "Unable to determine if collection exists. Not capping collection", "collection", colName)
		return false
	}

	if exists {
		klog.V(1).InfoS("Collection already exists. Capping could result in data loss. Ignoring", "collection", colName)
		return false
	}

	if strconv.IntSize < 64 {
		klog.V(1).InfoS("Pump running < 64bit architecture. Not capping collection as max size would be 2gb")
		return false
	}

	if colCapMaxSizeBytes == 0 {
		var defaultBytes int64 = 5
		colCapMaxSizeBytes = defaultBytes * 1024 * 1024 * 1024
		klog.InfoS("No max collection size set for connection, set default value", "connection", colName, "size", colCapMaxSizeBytes)
	}

	if colCapMaxDocuments == 0 {
		colCapMaxDocuments = 1000
		klog.InfoS("No max collection documents set for connection, set default value", "connection", colName, "size", colCapMaxDocuments)
	}

	cappedOptions := options.CreateCollection().
		SetCapped(true).
		SetSizeInBytes(colCapMaxSizeBytes).
		SetMaxDocuments(colCapMaxDocuments)

	err = ms.db.CreateCollection(context.Background(), ms.conf.CollectionName, cappedOptions)
	if err != nil {
		klog.ErrorS(err, "Unable to create capped collection", "collection", colName)
		return false
	}

	klog.InfoS("Capped collection created", "collection", colName, "bytes", colCapMaxSizeBytes, "docs", colCapMaxDocuments)

	return true
}

// collectionExists checks to see if a collection name exists in the db.
func (ms *MongoSink) collectionExists(name string) (bool, error) {
	colNames, err := ms.db.ListCollectionNames(context.Background(), bson.M{"name": name})
	if err != nil {
		klog.ErrorS(err, "Unable to get column names")
		return false, fmt.Errorf("failed to get collection name: %w", err)
	}

	for _, coll := range colNames {
		if coll == name {
			return true, nil
		}
	}

	return false, nil
}

// In returns an input channel for receiving data.
func (ks *MongoSink) In() chan<- any {
	return ks.in
}
