package Employee

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/labstack/echo"
	"github.com/radyatamaa/loyalti-go-echo/src/api/host/Config"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func PublishCreateEmployee(c echo.Context) error {
	//var data model.Merchant
	fmt.Println("masuk ke publisher")
	data := new(model.Employee)
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	//err := c.Bind(data)
	if err != nil {
		fmt.Println("Error Publisher : ", err.Error())
	}

	fmt.Println(data)

	//if len(data.MerchantEmail) == 0 && len(data.MerchantName) == 0 {
	//	return c.String(http.StatusBadRequest, "Nama dan Email kosong")
	//} else if len(data.MerchantEmail) == 0 {
	//	return c.String(http.StatusBadRequest, "Email tidak boleh kosong")
	//} else if len(data.MerchantName) == 0 {
	//	return c.String(http.StatusBadRequest, "Name tidak boleh kosong")
	//}

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

	var newTopic = "create-employee-topic"

	message, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error marshal : ", err.Error())
	}
	//message := `{
	//	"merchant_name":`+data.MerchantName+`,
	//	"merchant_email":"contact@jco.com"
	//}`
	msg := &sarama.ProducerMessage{
		Topic: newTopic,
		Value: sarama.StringEncoder(string(message)),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error bawah publisher : ", err.Error())
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", newTopic, partition, offset)
	return nil
}

