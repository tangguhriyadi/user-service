package service

import (
	"context"
	"errors"

	"github.com/tangguhriyadi/user-service/model"
	"github.com/tangguhriyadi/user-service/repository"
)

type UserService interface {
	GetAll(c context.Context) ([]model.Users, error)
	Create(c context.Context, userPayload *model.Users) error
	Update(c context.Context, userId string, payload *model.Users) error
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

	//check duplicat
	checkUsername, err := us.userRepo.FindByUsername(c, userPayload.Username)
	if checkUsername != nil {
		return errors.New("username is already exist")
	} else if err != nil {
		return err
	}

	//execute query create
	if err := us.userRepo.Create(c, userPayload); err != nil {
		return err
	}

	return nil
}

func (us UserServiceImpl) Update(c context.Context, userId string, payload *model.Users) error {

	if err := us.userRepo.Update(c, userId, payload); err != nil {
		return err
	}

	return nil

}
