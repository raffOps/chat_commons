package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/raffops/chat/internal/auth/repository"
	"github.com/raffops/chat/internal/errs"
	model "github.com/raffops/chat/internal/models"
	"github.com/raffops/chat/pkg/password_hasher"
	"net/http"
	"os"
	"time"
)

var (
	secretKey = []byte(os.Getenv("SECRET_KEY"))
)

type service struct {
	repo   repository.Repository
	hasher password_hasher.PasswordHasher
}

func (s *service) Login(ctx context.Context, name, password string) (string, *errs.Err) {
	user, err := s.repo.Get(ctx, name)
	if err != nil {
		return "", err
	}
	if !s.hasher.CheckPasswordHash(password, user.Password) {
		return "", &errs.Err{Message: "invalid password", Code: http.StatusUnauthorized}
	}

	claims := &model.Claims{
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Id,
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokeString, errSign := token.SignedString(secretKey)
	if errSign != nil {
		return "", &errs.Err{Message: "failed to sign token", Code: http.StatusInternalServerError}
	}
	return tokeString, nil
}

func (s *service) SignUp(ctx context.Context, name, password string) (model.User, *errs.Err) {
	user := model.User{Name: name, Password: password, Role: "USER"}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return model.User{}, &errs.Err{Message: err.Error(), Code: http.StatusBadRequest}
	}

	foundUser, errGet := s.repo.Get(ctx, user.Name)
	if errGet == nil && foundUser.Name == user.Name {
		return model.User{}, &errs.Err{Message: "user already exists", Code: http.StatusConflict}
	}

	user.Password, err = s.hasher.HashPassword(user.Password)
	if err != nil {
		return model.User{}, &errs.Err{Message: "failed to hash password", Code: http.StatusInternalServerError}
	}

	createdUser, errCreate := s.repo.Create(ctx, user)
	if errCreate != nil {
		return model.User{}, errCreate
	}
	return createdUser, nil
}

func NewUserService(repo repository.Repository, hasher password_hasher.PasswordHasher) Service {
	return &service{repo: repo, hasher: hasher}
}
