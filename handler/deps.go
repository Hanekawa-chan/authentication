package handler

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Generator struct {
	secretKey string
}

const UserID string = "user_id"

var (
	ErrInvalidToken error = errors.New("token not valid")
	ErrNotMapClaims error = errors.New("parsedToken.Claims not jwt.MapClaims")
	ErrIdNotFound   error = errors.New("id not found")
	ErrIsEmpty      error = errors.New("jwt field is empty")
)

func New(secretKey string) (*Generator, error) {
	return &Generator{secretKey: secretKey}, nil
}

func (g *Generator) Generate(claims map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims(claims))

	return token.SignedString([]byte(g.secretKey))
}

func (g *Generator) ParseToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(g.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrNotMapClaims
	}

	return claims, nil
}

func GetUserID(ctx context.Context) (uuid.UUID, error) {
	jwts := ctx.Value("jwt")
	if jwts == "" {
		return uuid.UUID{}, ErrIsEmpty
	}

	g := &Generator{}
	claims, err := g.ParseToken(jwts.(string))
	if err != nil {
		return uuid.UUID{}, err
	}

	id, ok := claims[UserID]
	if !ok {
		return uuid.UUID{}, ErrIdNotFound
	}

	userID, err := uuid.Parse(id.(string))
	if err != nil {
		return uuid.UUID{}, err
	}

	return userID, nil
}
