package controller

import (
	"net/http"

	"github.com/labstack/echo"
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
