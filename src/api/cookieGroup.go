package api

import (
    "github.com/radyatamaa/loyalti-go-echo/src/api/handlers"

    "github.com/labstack/echo"
)

func CookieGroup(g *echo.Group) {
    g.GET("/main", handlers.MainCookie)
}
