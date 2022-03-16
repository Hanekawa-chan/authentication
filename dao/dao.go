package dao

import (
	"authentication/models"
	"context"
	"errors"
)

type Auth interface {
	CreateRefresh(ctx context.Context, user *models.Credentials) error
}

var ErrNotFound = errors.New("not found")
var ErrAlreadyExists = errors.New("already exists")
