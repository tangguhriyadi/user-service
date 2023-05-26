package service

import (
	"context"
	"errors"

	"github.com/tangguhriyadi/user-service/dto"
	"github.com/tangguhriyadi/user-service/model"
	"github.com/tangguhriyadi/user-service/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAll(c context.Context, page int, limit int) (dto.AllUsers, error)
	Create(c context.Context, userPayload *model.Users) error
	Update(c context.Context, userId string, payload *model.Users) error
	GetById(c context.Context, userId string) (*model.Users, error)
	Delete(c context.Context, userId string) error
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (us UserServiceImpl) GetAll(c context.Context, page int, limit int) (dto.AllUsers, error) {
	result, err := us.userRepo.GetAll(c, page, limit)
	if err != nil {
		return dto.AllUsers{}, err
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPayload.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	userPayload.Password = string(hashedPassword)

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

func (us UserServiceImpl) GetById(c context.Context, userId string) (*model.Users, error) {
	result, err := us.userRepo.GetById(c, userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (us UserServiceImpl) Delete(c context.Context, userId string) error {
	if err := us.userRepo.Delete(c, userId); err != nil {
		return err
	}

	return nil
}
