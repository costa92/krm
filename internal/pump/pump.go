package pump

import (
	genericoptions "github.com/costa92/krm/pkg/options"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
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
