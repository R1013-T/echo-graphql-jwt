package main

import (
	"echo-graphql-jwt/database"
	"echo-graphql-jwt/routes"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	database.Init()
	e := routes.Init()

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
