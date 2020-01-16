// package middlewareCustom

// import (
// 	"net/http"
// 	"strconv"
// 	"sync"
// 	"time"

// 	"github.com/labstack/echo"
// )

// // middleware
// type (
// 	Stats struct {
// 		Uptime       time.Time      `json:"uptime"`
// 		RequestCount uint64         `json:"requestCount"`
// 		Statuses     map[string]int `json:"status"`
// 		mutex        sync.RWMutex
// 	}
// )

// func Newstats() *Stats {
// 	return &Stats{
// 		Uptime:   time.Now(),
// 		Statuses: map[string]int{},
// 	}
// }

// // process is for middleware
// func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		if err := next(c); err != nil {
// 			c.Error(err)
// 		}
// 		s.mutex.Lock()
// 		defer s.mutex.Unlock()
// 		s.RequestCount++
// 		status := strconv.Itoa(c.Response().Status)
// 		s.Statuses[status]++
// 		return nil
// 	}
// }

// // Handle is the endpoint to get stats
// func (s *Stats) Handle(c echo.Context) error {
// 	s.mutex.RLock()
// 	defer s.mutex.RUnlock()
// 	return c.JSON(http.StatusOK, s)
// }

// // ServerHeader middleware
// func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Response().Header().Set(echo.HeaderServer, "Echo/4.0")
// 		return next(c)
// 	}
// }
