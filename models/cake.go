package models

import (
	"TECHTEST_BE/database"
	"database/sql"
	"errors"
	"log"
	"time"
)

// Cake represents the cake object
type Cake struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Image       string  `json:"image"`
}

// CakeWithTimestamps represents
type CakeWithTimestamps struct {
	Cake
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	// DeletedAt string    `json:"deleted_at"`
}
var ErrCakeNotFound = errors.New("Cake not found")

// GetCakes retrieves the list of cakes from the database
func GetCakes() ([]Cake, error) {
	db := database.DB

	rows, err := db.Query("SELECT id, title, description, rating, image FROM cakes ORDER BY rating DESC, title ASC")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var cakes []Cake
	for rows.Next() {
		var cake Cake
		err := rows.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		cakes = append(cakes, cake)
	}

	return cakes, nil
}


// GetCake retrieves a specific cake by ID from the database
func GetCake(id int) (*CakeWithTimestamps, error) {
	db := database.DB

	var cake CakeWithTimestamps
	err := db.QueryRow("SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = ?", id).Scan(
		&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt,
	)
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

	currentTime := time.Now()

	_, err := db.Exec("INSERT INTO cakes (title, description, rating, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		cake.Title, cake.Description, cake.Rating, cake.Image, currentTime, currentTime)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// UpdateCake updates an existing cake in the database
func UpdateCake(id int, cake *Cake) error {
	db := database.DB

	currentTime := time.Now()

	result, err := db.Exec("UPDATE cakes SET title = ?, description = ?, rating = ?, image = ?, updated_at = ? WHERE id = ?",
		cake.Title, cake.Description, cake.Rating, cake.Image, currentTime, id)
	if err != nil {
		log.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		return ErrCakeNotFound
	}

	return nil
}


// DeleteCake deletes a cake from the database
func DeleteCake(id int) error {
	db := database.DB

	result, err := db.Exec("DELETE FROM cakes WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		return ErrCakeNotFound
	}

	return nil
}
