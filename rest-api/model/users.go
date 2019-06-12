package model

import (
	"database/sql"

	"github.com/lib/pq"
)

// Users struct
type Users struct {
	DB *sql.DB
}

// User struct
type User struct {
	ID        uint64
	Email     string
	Password  string
	CreatedAt string
}

// Login log in user
func (m Users) Login(user User) error {
	var hash string
	sqlStatement := `SELECT password FROM users WHERE email=$1`
	err := m.DB.QueryRow(sqlStatement, user.Email).Scan(&hash)

	if err == sql.ErrNoRows {
		return ErrUserNotFound
	}

	if err = checkPasswordHash(user.Password, hash); err != nil {
		return ErrIncorrectPass
	}
	return err
}

// Register  adds user to database
func (m Users) Register(user User) error {
	if !isValidPass(user.Password) {
		return ErrInvalidPass
	}
	if !validEmail.MatchString(user.Email) {
		return ErrInvalidEmail
	}
	hash, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	sqlStatement := `INSERT INTO users (email, password) VALUES($1, $2);`
	_, err = m.DB.Exec(sqlStatement, user.Email, hash)
	if err, ok := err.(*pq.Error); ok {
		if err.Code == "23505" {
			return ErrUserAlreadyExist
		}
	}
	return err
}
