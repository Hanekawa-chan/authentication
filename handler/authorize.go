package handler

import (
	"authentication/models"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Response struct {
	Jwt     string `json:"jwt_token,omitempty"`
	Refresh string `json:"refresh_token,omitempty"`
}

func (h *AuthHandler) Authorize(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	guid, err := uuid.Parse(id)
	if err != nil {
		log.Log().Err(err).Msg("uuid parse")
	}

	jwt, err := h.generateJWT(guid)
	if err != nil {
		log.Log().Err(err).Msg("generate jwt")
	}

	refresh, err := generateRefresh()
	err = h.db.CreateRefresh(context.Background(), &models.Credentials{Refresh: refresh})
	if err != nil {
		log.Log().Err(err).Msg("create refresh")
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

func (a *AuthHandler) generateJWT(userID uuid.UUID) (string, error) {
	claims := make(map[string]interface{})
	claims["user_id"] = userID
	token, err := a.generator.Generate(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateRefresh() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	hash := base64.URLEncoding.EncodeToString(b)

	return hash, nil
}
