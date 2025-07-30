package repository

import (
	"context"
	"template-service/internal/model"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*model.User, error)
}
