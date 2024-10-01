package mealplan

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbName = "meal_planner.db"
)

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	err = createTables()
	if err != nil {
		return fmt.Errorf("error creating tables: %w", err)
	}

	return nil
}

func createTables() error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS meals (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			day TEXT,
			meal_type TEXT,
			name TEXT
		)
	`)
	return err
}

func Save(mp *MealPlan) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM meals")
	if err != nil {
		return fmt.Errorf("error clearing existing meals: %w", err)
	}

	stmt, err := tx.Prepare("INSERT INTO meals (day, meal_type, name) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	for day, meals := range mp.Days {
		for _, meal := range meals {
			_, err := stmt.Exec(day, meal.Type, meal.Name)
			if err != nil {
				return fmt.Errorf("error inserting meal: %w", err)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	fmt.Println("Meal plan saved successfully!")
	return nil
}

func Load() *MealPlan {
	mp := New()

	rows, err := db.Query("SELECT day, meal_type, name FROM meals")
	if err != nil {
		log.Printf("Error querying meals: %v", err)
		return mp
	}
	defer rows.Close()

	for rows.Next() {
		var day, mealType, name string
		err := rows.Scan(&day, &mealType, &name)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		mp.AddMeal(day, mealType, name)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
	}

	fmt.Println("Meal plan loaded successfully!")
	return mp
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
