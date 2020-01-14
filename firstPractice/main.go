/**********************
* Golang Web Project
* Dev. Dongwon Paek
***********************/
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	collection := client.Database("test").Collection("trainers")
	fmt.Println(collection)
	// create instance
	e := echo.New()

	// use middleware for logging
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routing
	e.GET("/", MainPage)
	e.GET("/hello", hello)

	// Start Server
	e.Logger.Fatal(e.Start(":3000"))
}
