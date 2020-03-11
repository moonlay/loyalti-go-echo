package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/radyatamaa/loyalti-go-echo/src/api/host/Config"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/repository"
	"log"
	"os"
	"os/signal"
	"strings"
)

func consumeProgram(topics []string, master sarama.Consumer) (chan *sarama.ConsumerMessage, chan *sarama.ConsumerError) {
	consumers := make(chan *sarama.ConsumerMessage)
	errors := make(chan *sarama.ConsumerError)
	fmt.Println("Kafka Program is Ready")
	for _, topic := range topics {
		if strings.Contains(topic, "__consumer_offsets") {
			continue
		}
		partitions, _ := master.Partitions(topic)
		// this only consumes partition no 1, you would probably want to consume all partitions
		consumer, err := master.ConsumePartition(topic, partitions[0], sarama.OffsetNewest)
		if nil != err {
			fmt.Printf("Topic %v Partitions: %v", topic, partitions)
			panic(err)
		}
		//fmt.Println(" Start consuming topic ", topic)
		 go func(topic string, consumer sarama.PartitionConsumer) {
			for {
				select {
				case consumerError := <-consumer.Errors():
					errors <- consumerError
					fmt.Println("consumerError: ", consumerError.Err)

				case msg := <-consumer.Messages():
					//*messageCountStart++
					//Deserialize
					program := model.Program{}
					switch msg.Topic {
					case "create-program-topic":
						err := json.Unmarshal([]byte(msg.Value), &program)
						if err != nil {
							fmt.Println(err.Error())
							os.Exit(1)
						}
						repository.CreateProgram(&program)
						fmt.Println(string(msg.Value))
						fmt.Println("Berhasil membuat Program")
						break

					case "update-program-topic":
						err := json.Unmarshal([]byte(msg.Value), &program)
						if err != nil {
							fmt.Println(err.Error())
							os.Exit(1)
						}
						repository.UpdateProgram(&program)
						fmt.Println(string(msg.Value))
						fmt.Println("program berhasil di update")
						break

					case "delete-program-topic":
						err := json.Unmarshal([]byte(msg.Value), &program)
						if err != nil {
							fmt.Println(err.Error())
							os.Exit(1)
						}
						repository.DeleteProgram(&program)
					}
				}
			}
		}(topic, consumer)
	}

	return consumers, errors
}

func NewProgramConsumer() {

	brokers := []string{"11.11.5.146:9092"}

	kafkaConfig := Config.GetKafkaConfig("", "")

	master, err := sarama.NewConsumer(brokers, kafkaConfig)

	if err != nil {

		panic(err)

	}

	defer func() {

		if err := master.Close(); err != nil {

			panic(err)

		}

	}()

	//topic, err := master.Topics()
	if err != nil {
		panic(err)
	}
	topics, _ := master.Topics()
	//
	consumer, errors := consumeProgram(topics, master)
	////consumer1, err := master.ConsumePartition(updateTopic, 0, sarama.OffsetNewest)
	//
	if errors != nil {
		log.Println(errors)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Count how many message processed
	msgCount := 0

	// Get signnal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case msg := <-consumer:
				msgCount++
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
			case consumerError := <-errors:
				msgCount++
				fmt.Println("Received consumerError ", string(consumerError.Topic), string(consumerError.Partition), consumerError.Err)
				doneCh <- struct{}{}
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	master.Close()
	fmt.Println("Processed", msgCount, "messages")

}




