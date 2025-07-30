package repository

import (
	"context"
	"sample-service/internal/model"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*model.User, error)
}
