package api

import (
    "github.com/radyatamaa/loyalti-go-echo/src/api/handlers"
    "github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
    e.GET("/login", handlers.Login)
    e.GET("/yallo", handlers.Yallo)
    e.POST("/graphql", handlers.QuerySong)
}
