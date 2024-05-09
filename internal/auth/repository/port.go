package repository

import (
	"context"
	"github.com/raffops/chat/internal/errs"
	model "github.com/raffops/chat/internal/models"
)

//go:generate mockery --name Repository
type Repository interface {
	Get(ctx context.Context, name string) (model.User, *errs.Err)
	Create(ctx context.Context, user model.User) (model.User, *errs.Err)
}
