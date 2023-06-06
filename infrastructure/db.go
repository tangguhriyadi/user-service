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
		os.Getenv("DB_POSTGRES_PASSWORD"),
		os.Getenv("DB_POSTGRES_HOST"),
		os.Getenv("DB_POSTGRES_PORT"),
		os.Getenv("DB_POSTGRES_DBNAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Users{})
	DB = db
}
