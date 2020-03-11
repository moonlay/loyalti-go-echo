package main

import (
	"fmt"
	"github.com/radyatamaa/loyalti-go-echo/src/api/host/Config"
	"log"
	//"time"
	"github.com/Shopify/sarama"
)

//var (
//
//	brokerList = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:9092").Strings()
//
//	topic      = kingpin.Flag("topic", "Topic name").Default("important").String()
//
//	maxRetry   = kingpin.Flag("maxRetry", "Retry limit").Default("5").Int()
//
//)

func main() {

	//kingpin.Parse()
	//
	//config := sarama.NewConfig()
	//
	//config.Producer.RequiredAcks = sarama.WaitForAll
	//
	//config.Producer.Retry.Max = *maxRetry
	//
	//config.Producer.Return.Successes = true

	kafkaConfig := Config.GetKafkaConfig("", "")

	producer, err := sarama.NewSyncProducer([]string{"11.11.5.146:9092"}, kafkaConfig)

	if err != nil {

		panic(err)

	}

	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	var newTopic = "update-transaction-topic"

	message := `{
		"outlet_id":"2",
		"total_transaction":7000
	}`

	//message := `{
	//	"employee_name":"Felix",
	//	"employee_email":"felfolful10@gmail.com",
	//	"employee_pin":"123321",
	//	"outlet_id":"11"
	//}`
	// var updateTopic = "update-merchant-topic"
	//var createTopic = "new-outlet-topic"

	msg := &sarama.ProducerMessage{
		Topic: newTopic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", newTopic, partition, offset)
}

//func getKafkaConfig(username, password string) *sarama.Config {
//
//	kafkaConfig := sarama.NewConfig()
//
//	kafkaConfig.Producer.Return.Successes = true
//
//	kafkaConfig.Net.WriteTimeout = 5 * time.Second
//
//	kafkaConfig.Producer.Retry.Max = 0
//
//
//
//	if username != "" {
//
//		kafkaConfig.Net.SASL.Enable = true
//
//		kafkaConfig.Net.SASL.User = username
//
//		kafkaConfig.Net.SASL.Password = password
//
//	}
//
//	return kafkaConfig

//}
