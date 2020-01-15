/**********************
* Golang Web Project
* Dev. Dongwon Paek
***********************/
package main

import (
	"context"
	"fmt"
	"goWeb/handler"
	"log"
	"net/httpontext"

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

// DB connection
const connectionString = "mongodb://localhost:27017"

// DB name
const dbName = "test"

// Collection name
const collecName = "shoppingMall"

// collection object/instance
var collection *mongo.Collection

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func MainPage(c echo.Context) error {
	return c.String(http.StatusOK, "This is main page. Nothing else.")
}

// close mongodb
func closeDB(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func main() {
	// create instance
	e := echo.New()
	g := e.Group("/admin")

	// use middleware for logging
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: func(c echo.Context) bool {
			// Skip authentication for and signup login request
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	// connect to mongoDB
	defer closeDB(client)
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	collection := client.Database(dbName).Collection(collecName)
	fmt.Println(collection)

	// Init handler
	h := &handler.Handler{DB: db}

	// Routing
	e.GET("/", MainPage)
	e.GET("/hello", hello)
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
	e.POST("/follow/:id", h.Cart)
	e.GET("/product", h.Product)
	e.GET("/feed", h.FetchPost)

	// Start Server
	e.Logger.Fatal(e.Start(":3000"))
}
