package config

import (
	"backend-crs/model"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnv() {
	rootPath, err := filepath.Abs(".")
	if err != nil {
		panic("Cannot find the project root : " + err.Error())
	}
	envPath := filepath.Join(rootPath, ".env")

	err = godotenv.Load(envPath)
	if err != nil {
		panic("Error Loading the env file at " + envPath + ": " + err.Error())
	}
}

func GetDBURI() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
}

func ConnectDatabase() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(GetDBURI()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database " + err.Error())
	}

	err = db.AutoMigrate(&model.Student{}, &model.Admin{})
	if err != nil {
		panic("Failed to automigrate objects " + err.Error())
	}

	return db, err
}

func GetSecretKey() string {
	return os.Getenv("JWT_SECRET_KEY");
}