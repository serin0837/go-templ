package main

import (
	"github.com/labstack/echo/v4"
	"github.com/serin0837/go-templ/handler"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}
	app.GET("/", userHandler.HandlerUserShow)
	app.Start(":3333")

}

// frame work echo, fiber, Chi, Gin,n-> Echo func(context) error: he like echo package so that he can handle the error.
