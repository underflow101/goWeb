/**********************
* Golang Web Project
* Dev. Dongwon Paek
***********************/
package main

import (
	"net/http"
	"firstPractice/middlewareCustom"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func MainPage(c echo.Context) error {
	return c.String(http.StatusOK, "This is main page. Nothing else.")
}

func main() {
	// create instance
	e := echo.New()

	// use middleware for logging
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	s := middlewareCustom.NewStats()
	e.Use(s.Process)
	e.GET("/stats", s.Handle)

	e.Use(middlewareCustom.ServerHeader)

	// Routing
	e.GET("/", MainPage)
	e.GET("/hello", hello)

	// Start Server
	e.Logger.Fatal(e.Start(":3000"))
}
