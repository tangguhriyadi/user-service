package repository

import (
	"context"
	"errors"

	"github.com/tangguhriyadi/user-service/dto"
	"github.com/tangguhriyadi/user-service/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(c context.Context, page int, limit int) (dto.AllUsers, error)
	Create(c context.Context, payload *model.Users) error
	FindByUsername(c context.Context, username string) (*model.Users, error)
	Update(c context.Context, userId string, payload *model.Users) error
	GetById(c context.Context, userId string) (*model.Users, error)
	Delete(c context.Context, userId string) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (ur UserRepositoryImpl) GetAll(c context.Context, page int, limit int) (dto.AllUsers, error) {
	var userEntity []model.Users
	var count int64

	// var users []dto.UserData

	countResult := ur.db.WithContext(c).Model(&[]model.Users{}).Where("deleted =?", false).Count(&count)
	if countResult.Error != nil {
		return dto.AllUsers{}, countResult.Error
	}

	result := ur.db.WithContext(c).Select("full_name, email, age, religion, gender, photo, is_verified").Where("deleted =?", false).Offset((page - 1) * limit).Limit(limit).Find(&userEntity)
	if result.Error != nil {
		return dto.AllUsers{}, result.Error
	}

	// for _, v := range userEntity {
	// 	user := dto.UserData{
	// 		FullName:   v.FullName,
	// 		Email:      v.Email,
	// 		Age:        v.Age,
	// 		Religion:   v.Religion,
	// 		Gender:     v.Gender,
	// 		Photo:      v.Photo,
	// 		IsVerified: v.IsVerified,
	// 	}
	// 	users = append(users, user)
	// }

	var allUsers = dto.AllUsers{
		Data:       &userEntity,
		Page:       page,
		Limit:      limit,
		TotalItems: count,
	}

	return allUsers, nil
}

func (ur UserRepositoryImpl) Create(c context.Context, payload *model.Users) error {

	result := ur.db.WithContext(c).Create(&payload)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur UserRepositoryImpl) FindByUsername(c context.Context, username string) (*model.Users, error) {
	var user model.Users
	result := ur.db.WithContext(c).Where("username =?", username).First(&user).Find(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (ur UserRepositoryImpl) Update(c context.Context, userId string, payload *model.Users) error {
	result := ur.db.WithContext(c).Where("id = ?", userId).Updates(&payload)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur UserRepositoryImpl) GetById(c context.Context, userId string) (*model.Users, error) {
	var user model.Users
	result := ur.db.WithContext(c).First(&user, userId)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (ur UserRepositoryImpl) Delete(c context.Context, userId string) error {

	var user model.Users

	user.Deleted = true

	result := ur.db.WithContext(c).Where("id = ?", userId).Updates(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
