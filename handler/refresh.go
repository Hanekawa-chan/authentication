package handler

import (
	"authentication/models"
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *AuthHandler) Refresh(w http.ResponseWriter, req *http.Request) {
	token := mux.Vars(req)["token"]
	refresh, err := generateRefresh()
	id, err := h.db.ReplaceRefresh(context.Background(), &models.Credentials{Refresh: token},
		&models.Credentials{Refresh: refresh})
	if err != nil {
		log.Log().Err(err).Msg("replace refresh")
		WriteError(w, &ErrResponse{Err: err.Error()})
		return
	}

	guid, err := uuid.Parse(id)
	if err != nil {
		log.Log().Err(err).Msg("uuid parse")
		WriteError(w, &ErrResponse{Err: err.Error()})
		return
	}

	jwt, err := h.generateJWT(guid, refresh)
	if err != nil {
		log.Log().Err(err).Msg("generate jwt")
		WriteError(w, &ErrResponse{Err: err.Error()})
		return
	}

	response := &Response{
		Jwt:     jwt,
		Refresh: refresh,
	}

	Write(w, response)
}
