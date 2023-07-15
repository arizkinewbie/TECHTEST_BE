package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	DB = db
	log.Println("Database connected")
	return nil
}