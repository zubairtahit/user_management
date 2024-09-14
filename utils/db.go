package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"user_management/config"
)

// InitDB initializes the database connection and sets up connection pooling.
func InitDB() (*sql.DB, error) {
	dbConfig := config.LoadDBConfig()

	// Construct the connection string conditionally based on whether the password is provided
	var connStr string
	if dbConfig.Password != "" {
		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.SSLMode,
		)
	} else {
		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=%s",
			dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Name, dbConfig.SSLMode,
		)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error creating database connection: %w", err)
	}

	// Check if the connection is valid
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	// Configure connection pool settings.
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * 60)

	return db, nil
}
