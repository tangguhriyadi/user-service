package model

import (
	"time"
)

type Users struct { //entity
	Id         int        `gorm:"primaryKey" json:"-"`
	FullName   string     `gorm:"type:varchar(300)" json:"full_name" validate:"required"`
	Email      string     `gorm:"type:varchar(300)" json:"email" validate:"required"`
	Age        int        `gorm:"type:int" json:"age" validate:"required"`
	Religion   string     `gorm:"type:varchar(300)" json:"religion" validate:"required"`
	Gender     string     `gorm:"type:varchar(300)" json:"gender" validate:"required"`
	Username   string     `gorm:"type:varchar(300)" json:"username" validate:"required"`
	Password   string     `gorm:"type:varchar(300)" json:"password" validate:"required"`
	Photo      string     `gorm:"type:varchar(300)" json:"photo"`
	LastLogin  time.Time  `gorm:"type:timestamp;default:null" json:"last_login"`
	IsVerified bool       `gorm:"type:bool" json:"is_verified"`
	Deleted    bool       `gorm:"type:bool" json:"-"`
	DeletedAt  *time.Time `gorm:"type:timestamp;default:null" json:"-"`
}

type UserPayload struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Religion string `json:"religion"`
	Gender   string `json:"gender"`
	Username string `json:"username"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
}

type UserResponseCreate struct {
	Success bool `json:"success"`
}

type Login struct {
	Username string `gorm:"type:varchar(300)" json:"username" validate:"required"`
	Password string `gorm:"type:varchar(300)" json:"password" validate:"required"`
}

type LoginResponse struct {
	UserId int       `json:"user_id"`
	JWT    string    `json:"jwt"`
	Exp    time.Time `json:"exp"`
}
