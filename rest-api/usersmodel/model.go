package usersmodel

import "database/sql"

// Model struct
type Model struct {
	db *sql.DB
}

// User struct
type User struct {
	ID        uint64
	Email     string
	Password  string
	CreatedAt string
}

// New creates new model
func New(db *sql.DB) Model {
	return Model{db: db}
}

// Add adds user to database
func (m Model) Add(user User) (User, error) {
	sqlStatement := `INSERT INTO users (email, password) VALUES($1, $2) 
		RETURNING id, created_at`
	err := m.db.QueryRow(sqlStatement, user.Email, user.Password).Scan(
		&user.ID, &user.CreatedAt)
	return user, err
}

// Get gets user from database by email
func (m Model) Get(email string) (User, error) {
	user := User{}
	sqlStatement := `SELECT id, email, password, created_at FROM users WHERE email=$1`
	err := m.db.QueryRow(sqlStatement, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}
