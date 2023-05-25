package dto

import "time"

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
