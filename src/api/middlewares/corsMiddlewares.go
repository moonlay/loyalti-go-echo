   package middlewares

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetCorsMiddlewares(e *echo.Echo) {
	fmt.Println("masuk ke CORS")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{"GET", "POST"},
	}))
}
