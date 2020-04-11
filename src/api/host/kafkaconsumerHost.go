package host

import (
	"fmt"
	"github.com/radyatamaa/loyalti-go-echo/src/api/host/consumer"
)
func StartKafka() {
	fmt.Println("Kafka jalan dpianggil dari main 1")
	//consumer.NewReceiver()
	go consumer.NewMerchantConsumer()
	go consumer.NewOutletConsumer()
	go consumer.NewProgramConsumer()
	go consumer.NewEmployeeConsumer()
	go consumer.NewCardConsumer()
	go consumer.NewSpecialConsumer()
	go consumer.NewTransactionConsumer()
	go consumer.NewVoucherConsumer()
	go consumer.NewRewardConsumer()
}
