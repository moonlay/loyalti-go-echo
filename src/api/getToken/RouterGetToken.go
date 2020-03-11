package getToken

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func RouterGetToken(c echo.Context) error {
	//temp := model.ResponseCommand{}
	fmt.Println("masuk ke router")
	username := c.FormValue("username")
	password := c.FormValue("password")
	resp := GetToken(username, password)
	fmt.Println("Respon : ", *resp)

	return c.JSON(http.StatusOK, resp)
}
