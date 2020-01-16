package handler

import "gopkg.in/mgo.v2"

type (
	// DB Handler
	Handler struct {
		DB *mgo.Session
	}
)

const (
	Key = "secret"
)
