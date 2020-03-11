package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/radyatamaa/loyalti-go-echo/src/api/host/Config"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/repository"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

func consume(topics []string, master sarama.Consumer) (chan *sarama.ConsumerMessage, chan *sarama.ConsumerError) {
	consumers := make(chan *sarama.ConsumerMessage)
	errors := make(chan *sarama.ConsumerError)
	fmt.Println("Kafka Merchant is Ready")
	//fmt.Println(topics)
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
					merchant := model.NewMerchantCommand{}
					switch msg.Topic {
					case "create-merchant-topic":
						err := json.Unmarshal([]byte(msg.Value), &merchant)
						if err != nil {
							fmt.Println(err.Error())
							os.Exit(1)
						}
						fmt.Println("ini hasilnya : ",merchant)
						resp , err := repository.CreateMerchantWSO2(&merchant)
						repository.CreateMerchantWSO2(&merchant)
						fmt.Println("masuk ke fungsi WSO2")
						if err != nil {
							fmt.Println("Error :",err.Error())
							os.Exit(1)
						}
						//os.Exit(1)
						if resp.StatusCode != http.StatusCreated {
							bodyBytes, err := ioutil.ReadAll(resp.Body)
							if err != nil {
								fmt.Println("Error :",err.Error())
								log.Fatal(err)
								os.Exit(1)
							}
							bodyString := string(bodyBytes)
							fmt.Println(bodyString)
						}
						fmt.Println("ini respon")
						bodyBytes, err := ioutil.ReadAll(resp.Body)
						if err != nil {
							fmt.Println("Error :",err.Error())
							log.Fatal(err)
							os.Exit(1)
						}
						bodyString := string(bodyBytes)
						fmt.Println(bodyString)

						fmt.Println(string(msg.Value))
						fmt.Println("Berhasil Masuk Ke WSO2  asdas")
						fmt.Println("Merchant berhasil dibuat")
						break

					case "update-merchant-topic":
						err := json.Unmarshal([]byte(msg.Value), &merchant)
						if err != nil {
							fmt.Println(err.Error())
							os.Exit(1)
						}
						repository.UpdateMerchant2(&merchant)
						fmt.Println(string(msg.Value))
						fmt.Println("Merchant berhasil di-Update")

					case "delete-merchant-topic":
						err := json.Unmarshal([]byte(msg.Value), &merchant)
						if err != nil {
							fmt.Println(err.Error())
							os.Exit(1)
						}
						repository.DeleteMerchant2(&merchant)
						fmt.Println(string(msg.Value))
						fmt.Println("Merchant berhasil dihapus")
					}
				}
			}
		}(topic, consumer)
	}

	return consumers, errors
}

func NewMerchantConsumer() {

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
	consumer, errors := consume(topics, master)
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
