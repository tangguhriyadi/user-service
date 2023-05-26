package service

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tangguhriyadi/user-service/model"
	"github.com/tangguhriyadi/user-service/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(c context.Context, payload *model.Login) (*model.LoginResponse, error)
}

type AuthServiceImpl struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		userRepo: userRepo,
	}
}

var jwtKey = []byte("secret_key")

func (as AuthServiceImpl) Login(c context.Context, payload *model.Login) (*model.LoginResponse, error) {

	username, err := as.userRepo.FindByUsername(c, payload.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(username.Password), []byte(payload.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	var loginResponse model.LoginResponse

	loginResponse.UserId = username.Id
	loginResponse.Exp = time.Now().Add(time.Hour * 24)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = loginResponse.UserId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	loginResponse.JWT = tokenString

	return &loginResponse, nil
}
