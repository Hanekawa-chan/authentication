package main

import (
	"authentication/dao"
	"authentication/handler"
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

const uri = "mongodb://mongo:authpass@localhost:27017"
const port = ":8080"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal().Err(err).Msg("error connecting db")
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal().Err(err).Msg("error disconnecting db")
		}
	}()
	authDao := dao.New(client)

	h := handler.NewAuthHandler(authDao)

	log.Log().Msg("server started on " + port)
	err = http.ListenAndServe(port, h)
	if err != nil {
		log.Fatal().Err(err).Msg("error starting server")
	}
}
