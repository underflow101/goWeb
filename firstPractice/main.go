/**********************
* Golang Web Project
* Dev. Dongwon Paek
***********************/
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
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

	// create instance
	e := echo.New()

	// use middleware for logging
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routing
	e.Get("/hello", handle.MainPage())

	// Run Server
	e.Run(standard.New(":27017"))
}
