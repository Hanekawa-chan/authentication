package handler

import (
	"authentication/dao"
	"github.com/gorilla/mux"
)

type AuthHandler struct {
	*mux.Router
	db dao.Auth
}

func NewAuthHandler(db dao.Auth) *AuthHandler {
	r := mux.NewRouter()
	h := &AuthHandler{r, db}
	r.Queries("id", "{id}").Path("/auth").HandlerFunc(h.Authorize)
	r.Path("/refresh").HandlerFunc(h.Refresh)
	return h
}
