package handler

import (
	"authentication/models"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *AuthHandler) Refresh(w http.ResponseWriter, req *http.Request) {
	token := mux.Vars(req)["token"]
	refresh, err := generateRefresh()
	id, err = h.db.ReplaceRefresh(context.Background(), &models.Credentials{Refresh: token},
		&models.Credentials{Refresh: refresh})
	if err != nil {
		log.Log().Err(err).Msg("replace refresh")
	}

	jwt, err := h.generateJWT(guid, refresh)
	if err != nil {
		log.Log().Err(err).Msg("generate jwt")
	}

	response := &Response{
		Jwt:     jwt,
		Refresh: refresh,
	}

	res, err := json.Marshal(response)
	_, err = w.Write(res)
	if err != nil {
		log.Log().Err(err).Msg("write")
	}
}
