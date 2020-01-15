package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// DB Handler
	Handler struct {
		DB *mongo.Session
	}
)

const (
	Key = "secret"
)

// Systemic routes configuration
func RouteSetup(c *echo.Context) {
	// Ping Test
	c.GET("/ping", func(r *echo.Context) {
		r.String(http.StatusOK, "pong")
	})
}

func HandlerRoot(c *echo.Context) {
	c.JSON(http.StatusOK, echo.H{"healthcheck": "OK"})
}
