package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/high-la/movieapp/rating/pkg/model"
)

func main() {
	fmt.Println("Creating a Kafka producer")

	// created a Kafka producer by calling a kafka.NewProducer function and providing
	// localhost as the Kafka address for testing it locally.
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	const fileName = "ratingsdata.json"
	fmt.Println("Reading rating events from file " + fileName)

	// expected to read rating data from the ratingsdata.json file.
	ratingEvents, err := readRatingEvents(fileName)
	if err != nil {
		panic(err)
	}

	const topic = "ratings"
	if err := produceRatingEvents(topic, producer, ratingEvents); err != nil {
		panic(err)
	}

	const timeout = 10 * time.Second
	fmt.Println("Waiting " + timeout.String() + " until all events get produced")

	// Flush function to make sure all messages are sent to Kafka.
	producer.Flush(int(timeout.Milliseconds()))
}

func readRatingEvents(fileName string) ([]model.RatingEvent, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var ratings []model.RatingEvent
	if err := json.NewDecoder(f).Decode(&ratings); err != nil {
		return nil, err
	}
	return ratings, nil
}

func produceRatingEvents(topic string, producer *kafka.Producer, events []model.RatingEvent) error {
	for _, event := range events {
		encodedEvent, err := json.Marshal(event)
		if err != nil {
			return err
		}

		if err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny},
			Value: []byte(encodedEvent),
		}, nil); err != nil {
			return err
		}
	}
	return nil
}
