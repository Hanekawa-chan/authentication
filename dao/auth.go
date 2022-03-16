package dao

import (
	"authentication/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type auth struct {
	db *mongo.Client
}

func New(db *mongo.Client) Auth {
	return &auth{db: db}
}

func (a *auth) CreateRefresh(ctx context.Context, user *models.Credentials) error {
	const op = "dao.Auth.CreateRefresh"
	// добавить рефреш токен в бд
	return nil
}
