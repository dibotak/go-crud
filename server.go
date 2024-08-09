package main

import (
	"fmt"
	"go-crud/src/handler"
	router "go-crud/src/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	APP_PORT    string = "1323"
	DB_HOST     string = "localhost"
	DB_PORT     string = "5432"
	DB_USER     string = "postgres"
	DB_PASSWORD string = "postgres"
	DB_NAME     string = "gocrud"
)

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if os.Getenv("CUSTOM_ENV") != "TRUE" {
		return
	}
	APP_PORT = os.Getenv("APP_PORT")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
}

func main() {
	initEnv()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	handler.InitDB(db)

	r := router.New()
	r.Logger.Fatal(r.Start(":" + APP_PORT))
}
