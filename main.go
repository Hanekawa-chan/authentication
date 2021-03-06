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
		log.Fatal().Err(err).Msg("connect db")
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal().Err(err).Msg("disconnect db")
		}
	}()

	authDao := dao.New(client.Database("auth"))

	generator, err := handler.New("pass")
	if err != nil {
		log.Fatal().Err(err).Msg("create jwt generator")
	}

	h := handler.NewAuthHandler(authDao, generator)

	log.Log().Msg("server started on " + port)
	err = http.ListenAndServe(port, h)
	if err != nil {
		log.Fatal().Err(err).Msg("start server")
	}
}

//http://localhost:8080/auth?id=4428bdb1-f9b6-44bc-9467-f1463f32926c
//http://localhost:8080/refresh?refresh=
