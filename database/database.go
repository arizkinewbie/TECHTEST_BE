package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB merupakan objek database yang akan digunakan di seluruh aplikasi
var DB *sql.DB

// Inisialisasi koneksi ke database
func InitDB() error {
	db, err := sql.Open("mysql", "techtestbe:techtestbe@tcp(db4free.net:3306)/cakestoreralali")
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
