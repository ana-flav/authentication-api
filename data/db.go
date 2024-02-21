package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func DataBase() (*sql.DB, error) {
	
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err	)
		return nil, err
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil, err
	}

	return db, nil
}