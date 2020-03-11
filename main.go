package main

import (
	"fmt"

	"github.com/radyatamaa/loyalti-go-echo/src/router"
	//"github.com/spf13/viper"
)

// func init() {
// 	viper.SetConfigFile(`config.json`)
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// }
func main() {

	fmt.Println("Welcome to the webserver")
	e := router.New()
	//c := echo.New()
	//middlewares.SetCorsMiddlewares(c)

	// e.Start(viper.GetString("server.address"))

	//host.StartKafka()
	e.Start(":2525")
	//host.StartKafka()
	fmt.Println("Kafka start at port 2525")
}
