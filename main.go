package main

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

type authHandler struct {
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	method := parts[len(parts)-1]

	switch method {
	case "sign_up":
		// выдача access, refresh токенов по guid
	case "sign_in":
		// выдача refresh токена
	default:
		// выдача 404
	}
}

func main() {
	handler := &authHandler{}

	log.Log().Msg("server started on :8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal().Err(err).Msg("error starting server")
	}
}
