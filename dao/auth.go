package dao

import (
	"authentication/models"
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
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
	res := a.db.Collection("credentials").FindOne(ctx, bson.D{{"id", new.UserId}})
	var err error
	if res.Err() != nil {
		_, err = a.db.Collection("credentials").InsertOne(ctx, new)
	} else {
		return ErrAlreadyExists
	}
	return err
}

func (a *auth) ReplaceRefresh(ctx context.Context, last *models.Credentials, new *models.Credentials) (string, error) {
	// добавить рефреш токен в бд
	temp := &models.Credentials{}
	err := a.db.Collection("credentials").FindOne(ctx,
		bson.D{{"refresh", last.Refresh}}).Decode(temp)
	if err != nil {
		log.Log().Msg("no result")
		return "", err
	}
	id := temp.UserId
	new.UserId = id
	_, err = a.db.Collection("credentials").ReplaceOne(ctx, bson.D{{"refresh", last.Refresh}}, new)
	return id, err
}
