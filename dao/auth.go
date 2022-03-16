package dao

import "go.mongodb.org/mongo-driver/mongo"

type auth struct {
	db *mongo.Client
}

func New(db *mongo.Client) Auth {
	return &auth{db: db}
}
