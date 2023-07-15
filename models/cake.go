package models

import (
	"TECHTEST_BE/database"
	"database/sql"
	"log"
	"time"
)

// Cake merupakan representasi objek kue
type Cake struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Image       string  `json:"image"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   string  `json:"deleted_at"`
}

// GetCakes mengambil daftar kue dari database
func GetCakes() ([]Cake, error) {
	db := database.DB

	rows, err := db.Query("SELECT * FROM cakes ORDER BY rating DESC, title ASC")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var cakes []Cake
	for rows.Next() {
		var cake Cake
		err := rows.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt, &cake.DeletedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		cakes = append(cakes, cake)
	}

	return cakes, nil
}



// GetCake retrieves a specific cake by ID from the database
func GetCake(id int) (*Cake, error) {
		db := database.DB
	
		var cake Cake
		err := db.QueryRow("SELECT * FROM cakes WHERE id = ?", id).Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt, &cake.DeletedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil // Cake not found
			}
			return nil, err
		}
	
		return &cake, nil
	}

// AddCake adds a new cake to the database
func AddCake(cake *Cake) error {
	db := database.DB

	_, err := db.Exec("INSERT INTO cakes (title, description, rating, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		cake.Title, cake.Description, cake.Rating, cake.Image, time.Now(), time.Now())
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// UpdateCake updates an existing cake in the database
func UpdateCake(id int, cake *Cake) error {
	db := database.DB

	_, err := db.Exec("UPDATE cakes SET title = ?, description = ?, rating = ?, image = ?, updated_at = ? WHERE id = ?",
		cake.Title, cake.Description, cake.Rating, cake.Image, time.Now(), id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// DeleteCake deletes a cake from the database
func DeleteCake(id int) error {
	db := database.DB

	_, err := db.Exec("DELETE FROM cakes WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}