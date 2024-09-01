package pump

import (
	"github.com/costa92/krm/pkg/log"
	genericoptions "github.com/costa92/krm/pkg/options"
	kafkaconnector "github.com/costa92/krm/pkg/streams/connector/kafka"
	"github.com/segmentio/kafka-go"
	"github.com/superproj/onex/pkg/streams/flow"
	"go.mongodb.org/mongo-driver/mongo"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

type Config struct {
	KafkaOptions *genericoptions.KafkaOptions
	MongoOptions *genericoptions.MongoOptions
}

// Server contains state for a Kubernetes cluster master/api server.
type Server struct {
	config  kafka.ReaderConfig
	colName string
	db      *mongo.Database
}

// completedConfig is a completed Config. zh: completedConfig 是一个完成的 Config。
type completedConfig struct {
	*Config
}

// addUTC appends a UTC timestamp to the beginning of the message value.
var addUTC = func(msg kafka.Message) kafka.Message {
	timestamp := time.Now().Format(time.DateTime)

	// Concatenate the UTC timestamp with msg.Value
	msg.Value = []byte(timestamp + " " + string(msg.Value))
	return msg
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (cfg *Config) Complete() completedConfig {
	return completedConfig{cfg}
}

func (c *completedConfig) New() (*Server, error) {
	client, err := c.MongoOptions.NewClient()
	if err != nil {
		return nil, err
	}

	server := &Server{
		config: kafka.ReaderConfig{
			Brokers:           c.KafkaOptions.Brokers,
			Topic:             c.KafkaOptions.Topic,
			GroupID:           c.KafkaOptions.ReaderOptions.GroupID,
			QueueCapacity:     c.KafkaOptions.ReaderOptions.QueueCapacity,
			MinBytes:          c.KafkaOptions.ReaderOptions.MinBytes,
			MaxBytes:          c.KafkaOptions.ReaderOptions.MaxBytes,
			MaxWait:           c.KafkaOptions.ReaderOptions.MaxWait,
			ReadBatchTimeout:  c.KafkaOptions.ReaderOptions.ReadBatchTimeout,
			HeartbeatInterval: c.KafkaOptions.ReaderOptions.HeartbeatInterval,
			CommitInterval:    c.KafkaOptions.ReaderOptions.CommitInterval,
			RebalanceTimeout:  c.KafkaOptions.ReaderOptions.RebalanceTimeout,
			StartOffset:       c.KafkaOptions.ReaderOptions.StartOffset,
			MaxAttempts:       c.KafkaOptions.ReaderOptions.MaxAttempts,
		},
		colName: c.MongoOptions.Collection,
		db:      client.Database(c.MongoOptions.Database),
	}
	return server, nil
}

// preparedServer is a prepared Server. zh: preparedServer 是一个准备好的 Server。
type preparedServer struct {
	*Server
}

func (s preparedServer) Run(stopCh <-chan struct{}) error {
	ctx := wait.ContextForChannel(stopCh)

	source, err := kafkaconnector.NewKafkaSource(ctx, s.config)
	if err != nil {
		return err
	}
	filter := flow.NewMap(addUTC, 1)
	sink, err := mongoconnector.NewMongoSink(ctx, s.db, mongoconnector.SinkConfig{
		CollectionName:            s.colName,
		CollectionCapMaxDocuments: 2000,
		CollectionCapMaxSizeBytes: 5 * genericoptions.GiB,
		CollectionCapEnable:       true,
	})
	if err != nil {
		return err
	}

	log.Infof("Successfully start pump server")
	source.Via(filter).To(sink)

	return nil
}
