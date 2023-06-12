package infrastructure

import (
	"fmt"
	"os"

	"github.com/tangguhriyadi/user-service/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_POSTGRES_USER"),
		"user_service",
		os.Getenv("DB_POSTGRES_HOST"),
		"5433",
		"user_service",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Users{})
	DB = db
}
