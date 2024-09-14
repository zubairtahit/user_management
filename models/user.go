package models

import (
	"database/sql"
)

// UserRepository defines the methods for interacting with user data in the database.
type UserRepository interface {
	CreateUser(user User) (int, error)
	GetUser(id int) (User, error)
	UpdateUser(id int, user User) error
	DeleteUser(id int) error
	UserExists(id int) (bool, error)
	EmailExists(email string) (bool, error)
}

// DBInterface defines the methods for interacting with the database.
// This interface aids in mocking the database for testing.
type DBInterface interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// userRepositoryImpl implements UserRepository using a SQL database.
type userRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

// User represents a user model.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// CreateUser inserts a new user into the database and returns the new user ID.
func (r *userRepositoryImpl) CreateUser(user User) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO users(name, email, age) VALUES($1, $2, $3) RETURNING id", user.Name, user.Email, user.Age).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetUser retrieves a user by ID from the database.
func (r *userRepositoryImpl) GetUser(id int) (User, error) {
	var user User
	err := r.db.QueryRow("SELECT id, name, email, age FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err != nil {
		return user, err
	}
	return user, nil
}

// UpdateUser updates user details in the database.
func (r *userRepositoryImpl) UpdateUser(id int, user User) error {
	_, err := r.db.Exec("UPDATE users SET name = $1, email = $2, age = $3 WHERE id = $4", user.Name, user.Email, user.Age, id)
	return err
}

// DeleteUser removes a user from the database by ID.
func (r *userRepositoryImpl) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

// UserExists checks if a user with the given ID exists in the database.
func (r *userRepositoryImpl) UserExists(id int) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// EmailExists checks if an email already exists in the database.
func (r *userRepositoryImpl) EmailExists(email string) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
