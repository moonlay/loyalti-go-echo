package fcm

import (
	"encoding/json"
	"fmt"
	"github.com/NaySoftware/go-fcm"
	"github.com/labstack/echo"
)

const (
	serverKey = "AAAArkiu2Rs:APA91bFjMThIpOaeN5jkL-1EmyPQHrrKQ2mJ8S0uhl3tSHXBooSXZ_76Ht8QVxi0FplIR09ju3muQBlt2upMYe6xN7gv5RgZoWnlfLSVwl5ODqwC5yTLLG8LoTBiURMvjuXdqXaBcH4E"
)

type NotificationData struct {
	ReceiverToken string `json:"receiver_token"`
	Title string `json:"title"`
	Body string `json:"body"`
}

 func PushNotification(c echo.Context) error {
 	fmt.Println("masuk ke push notif")
 	var msg NotificationData
 	err := json.NewDecoder(c.Request().Body).Decode(&msg)
 	if err != nil {
 		fmt.Println("Error decode encode : ", err.Error())
	}
	data := map[string]string{
		"msg": msg.Title,
		"sum": msg.Body,
	}
	fmt.Println("data : ", data)

	ids := []string{
		msg.ReceiverToken,
	}

	fmt.Println("id token : ", ids)
	//xds := []string{
	//	msg.ReceiverToken,
	//}

	a := fcm.NewFcmClient(serverKey)
	a.NewFcmRegIdsMsg(ids, data)
	//c.AppendDevices(xds)

	status, err := a.Send()


	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}

	return nil
}

