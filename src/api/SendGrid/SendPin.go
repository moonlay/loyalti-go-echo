package SendGrid

import (
	//"encoding/json"
	"fmt"
	//"github.com/labstack/echo"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)


func SendPin(email model.EmailEmployee) error {
	fmt.Println("Masuk ke SendGrid")
	e := email
	//err := json.NewDecoder(c.Request().Body).Decode(&e)
	//if err != nil {
	//	fmt.Println("Error SendGrid : ", err.Error())
	//}

	//var e Email
	var a = "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"

	from := mail.NewEmail(e.SenderName, e.SenderEmail)
	subject := e.Subject
		fmt.Println("masuk ke perulangan")
		to := mail.NewEmail(e.EmployeeName, e.EmployeeEmail)
		plainTextContent := "abcd"
		htmlContent := "Hi , " + e.EmployeeName + "You've been invited to " + "<b>LoyaltiExpress</b>." + "This is yout PIN : " + e.EmployeePin
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
	fmt.Println("Selesai Tanpa Error : ", e)
	return nil
}
