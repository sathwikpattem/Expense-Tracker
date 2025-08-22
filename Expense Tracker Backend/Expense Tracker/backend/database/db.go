package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the database connection and creates tables
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./expenses.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	createTables()
	insertDefaultCategories()
	log.Println("Database initialized successfully")
}

// createTables creates the necessary database tables
func createTables() {
	// Create categories table
	createCategoriesTable := `
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL
	);`

	// Create expenses table
	createExpensesTable := `
	CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL NOT NULL,
		category TEXT NOT NULL,
		date TEXT NOT NULL,
		note TEXT NOT NULL,
		created DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (category) REFERENCES categories(name)
	);`

	if _, err := DB.Exec(createCategoriesTable); err != nil {
		log.Fatal("Failed to create categories table:", err)
	}

	if _, err := DB.Exec(createExpensesTable); err != nil {
		log.Fatal("Failed to create expenses table:", err)
	}
}

// insertDefaultCategories inserts default categories if they don't exist
func insertDefaultCategories() {
	defaultCategories := []string{"Food", "Travel", "Groceries", "Entertainment", "Others"}
	
	for _, category := range defaultCategories {
		_, err := DB.Exec("INSERT OR IGNORE INTO categories (name) VALUES (?)", category)
		if err != nil {
			log.Printf("Failed to insert default category %s: %v", category, err)
		}
	}
}
