package service

import (
	"context"
	"github.com/raffops/chat/internal/errs"
	model "github.com/raffops/chat/internal/models"
)

//go:generate mockery --name Service
type Service interface {
	Login(ctx context.Context, name, password string) (string, *errs.Err)
	SignUp(ctx context.Context, user model.User) (model.User, *errs.Err)
}
