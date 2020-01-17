package handler

import "go.mongodb.org/mongo-driver/mongo"

type (
	// DB Handler
	Handler struct {
		DB *mongo.Client
	}
)

const (
	Key = "secret"
)
