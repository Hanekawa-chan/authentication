package dao

import (
	"authentication/models"
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type auth struct {
	db *mongo.Database
}

func New(db *mongo.Database) Auth {
	return &auth{db: db}
}

func (a *auth) CreateRefresh(ctx context.Context, new *models.Credentials) error {
	// добавить рефреш токен в бд
	_, err := a.db.Collection("credentials").InsertOne(ctx, new)
	return err
}

func (a *auth) ReplaceRefresh(ctx context.Context, last *models.Credentials, new *models.Credentials) (string, error) {
	// добавить рефреш токен в бд
	res := a.db.Collection("credentials").FindOne(ctx, last)
	if res.Err() != nil {
		return "", res.Err()
	}
	id, err := res.DecodeBytes()
	log.Log().Msg(id.String())
	_, err := a.db.Collection("credentials").ReplaceOne(ctx, last, new)
	if err != nil {
		return "", err
	}
	_, err = a.db.Collection("credentials").InsertOne(ctx, new)
	return id, err
}
