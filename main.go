package main

import (
	"authentication/handler"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.Queries("id", "{id}").Path("").HandlerFunc(handler.Authorize)

	log.Log().Msg("server started on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal().Err(err).Msg("error starting server")
	}
}
