package dto

import (
	"time"

	"github.com/tangguhriyadi/user-service/model"
)

type Users struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Religion string `json:"religion"`
	Gender   string `json:"gender"`
	Username string `json:"username"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
}

type AllUsers struct {
	Data       *[]model.Users `json:"data"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalItems int64          `json:"total_items"`
}

type UserUpdate struct {
	FullName   string     `gorm:"type:varchar(300)" json:"full_name" `
	Email      string     `gorm:"type:varchar(300)" json:"email" `
	Age        int        `gorm:"type:int" json:"age" `
	Religion   string     `gorm:"type:varchar(300)" json:"religion" `
	Gender     string     `gorm:"type:varchar(300)" json:"gender" `
	Username   string     `gorm:"type:varchar(300)" json:"username" `
	Password   string     `gorm:"type:varchar(300)" json:"password" `
	Photo      string     `gorm:"type:varchar(300)" json:"photo"`
	IsVerified bool       `gorm:"type:bool" json:"is_verified"`
	Deleted    bool       `gorm:"type:bool" json:"deleted"`
	DeletedAt  *time.Time `gorm:"type:timestamp;default:null" json:"deleted_at"`
}

type UserParams struct {
	Id int `gorm:"type:int" uri:"id" validate:"required"`
}

type AllUsersParams struct {
	page  int `gorm:"type:int;default:1" json:"limit"`
	limit int `gorm:"type:int;default:10" json:"total_limit"`
}

type UserData struct {
	FullName   string `gorm:"type:varchar(300)" json:"full_name" `
	Email      string `gorm:"type:varchar(300)" json:"email" `
	Age        int    `gorm:"type:int" json:"age" `
	Religion   string `gorm:"type:varchar(300)" json:"religion" `
	Gender     string `gorm:"type:varchar(300)" json:"gender" `
	Photo      string `gorm:"type:varchar(300)" json:"photo"`
	IsVerified bool   `gorm:"type:bool" json:"is_verified"`
}
