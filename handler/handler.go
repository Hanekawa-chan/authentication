package handler

import (
	"authentication/dao"
	"github.com/gorilla/mux"
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
