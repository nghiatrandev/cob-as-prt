package dataservice

import (
	"fmt"
	. "github.com/nghiatrandev/sample_project/common"
	"github.com/segmentio/kafka-go"
	"net"
	"strconv"
)

type Kafka struct {
	topics  Topics
	writers map[string]*kafka.Writer
	readers map[string]*kafka.Reader
}

func NewKafka(config KafkaConfig) *Kafka {

	// Build a reader map of reader name -> kafka reader object
	readTopics := []string{config.Topics.TestTopic}
	readers := make(map[string]*kafka.Reader)
	for _, topic := range readTopics {
		topicName := topic
		readers[topicName] = kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{config.Addr},
			Topic:    topicName,
			GroupID:  "group-" + topicName,
			MinBytes: 10e6,
			MaxBytes: 10e6,
		})
	}

	// Build a writer map of writer name -> kafka writer object
	writeTopics := []string{config.Topics.TestTopic, config.Topics.TestTopic}
	writers := make(map[string]*kafka.Writer)
	for _, topic := range writeTopics {
		topicName := topic
		writers[topicName] = &kafka.Writer{
			Addr:         kafka.TCP(config.Addr),
			Topic:        topicName,
			Balancer:     &kafka.LeastBytes{},
			RequiredAcks: kafka.RequireAll,
			BatchSize:    1,
		}
	}

	// Create all topics (will do nothing if the topics had been already created)
	if err := createTopics(config, append(writeTopics, readTopics...)); err != nil {
		fmt.Errorf("failed to create topics: %s", err)
	}

	return &Kafka{
		writers: writers,
		readers: readers,
		topics:  config.Topics,
	}
}

//func (k *Kafka) publish(ctx context.Context, topic string, data []byte) error {
//	if writer, ok := k.writers[topic]; !ok {
//		return fmt.Errorf("there is no writer registered for topic %s", topic)
//	} else {
//		return writer.WriteMessages(ctx, kafka.Message{
//			Value: data,
//		})
//	}
//}
//
//func (k *Kafka) consume(ctx context.Context, topic string, handler func(context.Context, []byte) error) error {
//	reader, ok := k.readers[topic]
//	if !ok {
//		return fmt.Errorf("there is no reader registered for topic %s", topic)
//	}
//
//	msg, err := reader.ReadMessage(ctx)
//	if err != nil {
//		return err
//	}
//
//	return handler(ctx, msg.Value)
//}

func createTopics(config KafkaConfig, topicNames []string) error {
	conn, err := kafka.Dial("tcp", config.Addr)
	if err != nil {
		return err
	}
	controller, err := conn.Controller()
	if err != nil {
		return err
	}
	connController, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return err
	}

	for _, topic := range topicNames {
		// TODO config number of partitions & replica here
		if err := connController.CreateTopics(kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     -1,
			ReplicationFactor: -1,
		}); err != nil {
			return err
		}
	}

	return nil
}
