package handler

import (
	"authentication/dao"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

type AuthHandler struct {
	*mux.Router
	db        dao.Auth
	generator *Generator
}

func NewAuthHandler(db dao.Auth, gen *Generator) *AuthHandler {
	r := mux.NewRouter()
	h := &AuthHandler{r, db, gen}
	r.Queries("id", "{id}").Path("/auth").HandlerFunc(h.Authorize)
	r.Queries("token", "{token}").Path("/refresh").HandlerFunc(h.Refresh)
	return h
}

type Response struct {
	Jwt     string `json:"jwt_token,omitempty"`
	Refresh string `json:"refresh_token,omitempty"`
}

type ErrResponse struct {
	Err string `json:"error,omitempty"`
}

func Write(w http.ResponseWriter, response *Response) {
	res, err := json.Marshal(response)
	_, err = w.Write(res)
	if err != nil {
		log.Log().Err(err).Msg("write")
	}
}

func WriteError(w http.ResponseWriter, response *ErrResponse) {
	res, err := json.Marshal(response)
	_, err = w.Write(res)
	if err != nil {
		log.Log().Err(err).Msg("write")
	}
}
