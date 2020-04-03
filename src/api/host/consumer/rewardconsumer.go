package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/radyatamaa/loyalti-go-echo/src/api/host/Config"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/repository"
	"os"
	"os/signal"
	"strings"
)

func consumeReward(topics []string, master sarama.Consumer) (chan *sarama.ConsumerMessage, chan *sarama.ConsumerError) {
	consumers := make(chan *sarama.ConsumerMessage)
	errors := make(chan *sarama.ConsumerError)
	fmt.Println("kafka Reward is ready")
	for _, topic := range topics {
		if strings.Contains(topic, "__consumer_offsets") {
			continue
		}
		partitions, _ := master.Partitions(topic)
		// this only consumes partition no 1, you would probably want to consume all partitions
		consumer, err := master.ConsumePartition(topic, partitions[0], sarama.OffsetNewest)
		if nil != err {
			fmt.Println("Erorr reward : ", err.Error())
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
					reward := model.Reward{}
					switch msg.Topic {
					case "create-reward-topic":
						err := json.Unmarshal([]byte(msg.Value), &reward)
						if err != nil {
							fmt.Println(err.Error())
							os.Exit(1)
						}
						repository.CreateReward(&reward)
						fmt.Println(string(msg.Value))
						fmt.Println("Reward berhasil dibuat")
						break
					case "update-reward-topic":
						err := json.Unmarshal([]byte(msg.Value), &reward)
						if err != nil {
							fmt.Println(err.Error())
							os.Exit(1)
						}
						repository.UpdateReward(&reward)
						fmt.Println(string(msg.Value))
						fmt.Println("reward berhasil diupdate")
						break
					case "delete-reward-topic":
						err := json.Unmarshal([]byte(msg.Value), &reward)
						if err != nil {
							fmt.Println(err.Error())
							os.Exit(1)
						}
						repository.DeleteReward(&reward)
						fmt.Println("reward berhasil dihapus")
						break
					}
				}
			}
		}(topic, consumer)
	}

	return consumers, errors
}

func NewRewardConsumer() {

	brokers := []string{"52.185.161.109:9092"}
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
	consumer, errors := consumeReward(topics, master)
	////consumer1, err := master.ConsumePartition(updateTopic, 0, sarama.OffsetNewest)
	//
	if errors != nil {
		fmt.Println(err)
		//panic(err)

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

