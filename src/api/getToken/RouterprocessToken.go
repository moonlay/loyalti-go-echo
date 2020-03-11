package getToken

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func RouterProcessToken (c echo.Context) error{
	fmt.Println("masuk ke router proces Token")
	token := c.Request().Header.Get("Authorization")
	response := ProcessToken(token)
	fmt.Println("response : ", response)

	return c.JSON(http.StatusOK, response)
}
