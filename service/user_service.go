package service

import (
	"context"

	"github.com/tangguhriyadi/user-service/model"
	"github.com/tangguhriyadi/user-service/repository"
)

type UserService interface {
	GetAll(c context.Context) ([]model.Users, error)
	Create(c context.Context, userPayload *model.Users) error
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (us UserServiceImpl) GetAll(c context.Context) ([]model.Users, error) {
	result, err := us.userRepo.GetAll(c)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (us UserServiceImpl) Create(c context.Context, userPayload *model.Users) error {
	err := us.userRepo.Create(c, userPayload)
	if err != nil {
		return err
	}
	return nil
}
