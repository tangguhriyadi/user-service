package model

import (
	"time"
)

type Users struct { //entity
	Id         int        `gorm:"primaryKey" json:"id"`
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
	Deleted    bool       `gorm:"type:bool" json:"deleted"`
	DeletedAt  *time.Time `gorm:"type:timestamp;default:null" json:"deleted_at"`
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
