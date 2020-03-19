package SendGrid

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

)

type Email struct {
	//Sender      string   `json:"sender"`
	Receiver    []ReceiverStruct `json:"receiver"`
	Subject     string           `json:"subject"`
	Body        string           `json:"body"`
	TextContent string           `json:"text_content"`
}

type ReceiverStruct struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}


func SendMail(c echo.Context) error{
	fmt.Println("Masuk ke SendGrid")
	var e Email
	err := json.NewDecoder(c.Request().Body).Decode(&e)
	if err != nil {
		fmt.Println("Error SendGrid : ", err.Error())
	}

	//var e Email
	var a = "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"

	from := mail.NewEmail("LOYALTIExpress", "felixsiburian10@gmail.com")
	subject := e.Subject
	for i := range e.Receiver {
		fmt.Println("masuk ke perulangan")
		to := mail.NewEmail(e.Receiver[i].Name, e.Receiver[i].Email)
		plainTextContent := "abcd"
		htmlContent := e.Body
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client := sendgrid.NewSendClient(a)
		fmt.Println("isi sendgrid ", a)
		response, err := client.Send(message)
		if err != nil {
			fmt.Println("Error response : ", err.Error())
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
		}

		if err == nil {
			fmt.Println("mail sent")
		}
	}
	fmt.Println("Selesai Tanpa Error : ", e)
	return nil
}
