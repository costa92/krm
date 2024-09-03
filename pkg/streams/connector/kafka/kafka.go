package kafka

import (
	"context"
	"github.com/costa92/krm/pkg/streams"
	"github.com/costa92/krm/pkg/streams/flow"
	"github.com/segmentio/kafka-go"
	"k8s.io/klog/v2"
	"os"
	"os/signal"
	"syscall"
)

// 参考: https://gitcode.com/gh_mirrors/go/go-streams/blob/master/kafka/kafka_sarama.go
// KafkaSource represents an Apache Kafka source connector.
type KafkaSource struct {
	r         *kafka.Reader
	out       chan any
	ctx       context.Context
	cancelCtx context.CancelFunc
}

// NewKafkaSource returns a new KafkaSource instance.
func NewKafkaSource(ctx context.Context, config kafka.ReaderConfig) (*KafkaSource, error) {
	out := make(chan any)
	cctx, cancel := context.WithCancel(ctx)

	sink := &KafkaSource{
		r:         kafka.NewReader(config),
		out:       out,
		ctx:       cctx,
		cancelCtx: cancel,
	}

	go sink.init()
	return sink, nil
}

// init starts the main loop.
func (ks *KafkaSource) init() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	go ks.consume()

	select {
	case <-sigchan:
		ks.cancelCtx()
	case <-ks.ctx.Done():
	}

	close(ks.out)
	ks.r.Close()
}

// consume reads messages from the Kafka topic and sends them to the output channel.
//// zh: consume 从 Kafka 主题中读取消息并将其发送到输出通道。
//func (ks *KafkaSource) consume() {
//	for {
//		msg, err := ks.r.FetchMessage(ks.ctx)
//		if err != nil {
//			klog.Errorf("error while fetching message: %v", err)
//			break
//		}
//		fmt.Println("message received: ", string(msg.Value))
//		ks.out <- msg.Value
//		ks.r.CommitMessages(ks.ctx, msg)
//	}
//}

// consume reads messages from the Kafka topic and sends them to the output channel.
// zh: consume 从 Kafka 主题中读取消息并将其发送到输出通道。
func (ks *KafkaSource) consume() {
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := ks.r.ReadMessage(ks.ctx)
		if err != nil {
			klog.ErrorS(err, "Failed to read message")
		}
		ks.out <- msg
	}
}

// Via streams data through the given flow.
// zh: 通过给定的流传输数据。
func (ks *KafkaSource) Via(_flow streams.Flow) streams.Flow {
	flow.DoStream(ks, _flow)
	return _flow
}

// Out returns an output channel for sending data.
// zh: Out 返回一个用于发送数据的输出通道。
func (ks *KafkaSource) Out() <-chan any {
	return ks.out
}

// KafkaSink represents an Apache Kafka sink connector.
// zh: KafkaSink 表示 Apache Kafka 接收器连接器。
type KafkaSink struct {
	ctx context.Context
	w   *kafka.Writer
	in  chan any
}

// NewKafkaSink returns a new KafkaSink instance.
// zh: NewKafkaSink 返回一个新的 KafkaSink 实例。
func NewKafkaSink(ctx context.Context, config kafka.WriterConfig) (*KafkaSink, error) {
	sink := &KafkaSink{
		ctx: ctx,
		w:   kafka.NewWriter(config),
		in:  make(chan any),
	}

	go sink.init()
	return sink, nil
}

// init starts the main loop.
// zh: init 启动主循环。
func (ks *KafkaSink) init() {
	for msg := range ks.in {
		var km kafka.Message
		switch m := msg.(type) {
		case []byte:
			km.Value = m
		case string:
			km.Value = []byte(m)
		case *kafka.Message:
			km = *m
		default:
			klog.V(1).InfoS("Unsupported message type", "message", m)
			continue
		}
		if err := ks.w.WriteMessages(ks.ctx, km); err != nil {
			klog.ErrorS(err, "Failed to write message")
		}
	}

	ks.w.Close()
}

// In returns an input channel for receiving data.
// zh: In 返回一个用于接收数据的输入通道。
func (ks *KafkaSink) In() chan<- any {
	return ks.in
}
