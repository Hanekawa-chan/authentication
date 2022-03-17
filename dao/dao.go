package dao

import (
	"authentication/models"
	"context"
	"errors"
)

type Auth interface {
	CreateRefresh(ctx context.Context, new *models.Credentials) error
	ReplaceRefresh(ctx context.Context, last *models.Credentials, new *models.Credentials) (string, error)
}

var ErrNotFound = errors.New("not found")
var ErrAlreadyExists = errors.New("already exists")
