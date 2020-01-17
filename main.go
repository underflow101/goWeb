/**********************
* Golang Web Project
* Dev. Dongwon Paek
***********************/
package main

import (
	"context"
	"fmt"
	"goWeb/handler"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name     string
	Password string
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
	e.Logger.SetLevel(log.ERROR)

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
	// db, err := mgo.Dial("mongo")
	// // if err != nil {
	// // 	e.Logger.Fatal(err)
	// // }

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	db, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB! :)")

	// create index
	// if err = db.Copy().DB("bird").C("users").EnsureIndex(mgo.Index{
	// 	Key:    []string{"email"},
	// 	Unique: true,
	// }); err != nil {
	// 	log.Fatal(err)
	// }
	collection := db.Database("test").Collection("trainers")

	bear := Trainer{"Bear", "12345"}

	insertRes, err := collection.InsertOne(context.TODO(), bear)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single user: ", insertRes.InsertedID)

	// Init handler
	h := &handler.Handler{DB: db}

	// Routing
	e.GET("/", MainPage)
	e.GET("/hello", hello)
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
	e.POST("/follow/:id", h.Follow)
	//e.GET("/product", h.Product)
	e.GET("/feed", h.FetchPost)
	e.POST("/posts", h.CreatePost)

	// Start Server
	e.Logger.Fatal(e.Start(":3000"))
}
