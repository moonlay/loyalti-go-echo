package router

import (
	"github.com/labstack/echo"
	"github.com/radyatamaa/loyalti-go-echo/src/api"
	"github.com/radyatamaa/loyalti-go-echo/src/api/SendGrid"
	"github.com/radyatamaa/loyalti-go-echo/src/api/apiKafka/Card"
	"github.com/radyatamaa/loyalti-go-echo/src/api/apiKafka/Employee"
	"github.com/radyatamaa/loyalti-go-echo/src/api/apiKafka/Merchant"
	"github.com/radyatamaa/loyalti-go-echo/src/api/apiKafka/Outlet"
	"github.com/radyatamaa/loyalti-go-echo/src/api/apiKafka/Program"
	"github.com/radyatamaa/loyalti-go-echo/src/api/apiKafka/SpecialProgram"
	"github.com/radyatamaa/loyalti-go-echo/src/api/apiKafka/TransactionMerchant"
	"github.com/radyatamaa/loyalti-go-echo/src/api/apiKafka/Voucher"
	"github.com/radyatamaa/loyalti-go-echo/src/api/apiKafka/Reward"
	"github.com/radyatamaa/loyalti-go-echo/src/api/fcm"

	//"github.com/radyatamaa/loyalti-go-echo/src/api/fcm"

	//"github.com/radyatamaa/loyalti-go-echo/src/api/fcm"
	"github.com/radyatamaa/loyalti-go-echo/src/api/getToken"
	"github.com/radyatamaa/loyalti-go-echo/src/api/middlewares"
	"net/http"
)

func New() *echo.Echo {
	// create groups
	e := echo.New()
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	// set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetCookieMiddlewares(cookieGroup)
	middlewares.SetJwtMiddlewares(jwtGroup)
	middlewares.SetCorsMiddlewares(e)

	// set main routes
	api.MainGroup(e)

	// set group routes
	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGroup)
	api.JwtGroup(jwtGroup)

	e.GET("/ping", Ping)
	//host.StartKafka()

	//Kafka Merchant
	e.POST("/create-merchant", Merchant.PublishCreateMerchant)
	e.POST("/update-merchant", Merchant.PublishUpdateMerchant)
	e.POST("/delete-merchant", Merchant.PublishDeleteMerchant)

	//Kafka Card
	e.POST("/create-card", Card.PublishCreateCard)
	e.POST("/update-card", Card.PublishUpdateCard)
	e.POST("/delete-card", Card.PublishDeleteCard)

	//Kafka Employee
	e.POST("/create-employee", Employee.PublishCreateEmployee)
	e.POST("/update-employee", Employee.PublishUpdateEmployee)
	e.POST("/delete-employee", Employee.PublishDeleteEmployee)

	//Kafka Outlet
	e.POST("/create-outlet", Outlet.PublishCreateOutlet)
	e.POST("update-outlet", Outlet.PublishUpdateOutlet)
	e.POST("/delete-outlet", Outlet.PublishDeleteOutlet)

	//Kafka Program
	e.POST("/create-program", Program.PublishCreateProgram)
	e.POST("/update-program ", Program.PublishUpdateProgram)
	e.POST("/delete-program", Program.PublishDeleteProgram)

	//Kafka Special Program
	e.POST("/create-special", SpecialProgram.PublishCreateSpecial)
	e.POST("/update-special", SpecialProgram.PublishUpdateSpecial)
	e.POST("/delete-special", SpecialProgram.PublishDeleteSpecial)

	//Kafka Transaction Merchant
	e.POST("/create-transaction", TransactionMerchant.PublishCreateTransaction)
	e.POST("/update-transaction", TransactionMerchant.PublishUpdateTransaction)
	e.POST("/delete-transaction", TransactionMerchant.PublishDeleteTransaction)

	//Kafka Voucher Merchant
	e.POST("/create-voucher", Voucher.PublishCreateVoucher)
	e.POST("/update-voucher", Voucher.PublishUpdateVoucher)
	e.POST("/delete-voucher", Voucher.PublishDeleteTransaction)

	e.POST("/create-reward", Reward.PublishCreateReward)
	e.POST("/update-reward", Reward.PublishUpdateReward)
	e.POST("/delete-reward", Reward.PublishUpdateReward)

	//Get Token
	e.POST("/getToken", getToken.RouterGetToken)
	e.GET("/processToken", getToken.RouterProcessToken)

	//Post FCM
	e.POST("/getFCM", fcm.PushNotification)

	//Send Mail SendGrid
	e.POST("/SendMail", SendGrid.SendMail)

	return e
}

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

