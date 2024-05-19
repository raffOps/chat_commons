package service

import (
	"context"
	"github.com/raffops/chat/internal/errs"
	model "github.com/raffops/chat/internal/models"
)

type Service interface {
	Login(ctx context.Context, name, password string) (string, *errs.Err)
	SignUp(ctx context.Context, name, password string) (model.User, *errs.Err)
}
