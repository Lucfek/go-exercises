package model

import (
	"database/sql"

	"github.com/lib/pq"
)

// User struct
type User struct {
	ID        uint64
	Email     string
	Password  string
	CreatedAt string
}

// Add adds user to database
func (m Model) Login(user User) (User, error) {
	var hash string
	sqlStatement := `SELECT id, email, password, created_at FROM users WHERE email=$1`
	err := m.db.QueryRow(sqlStatement, user.Email).Scan(&user.ID, &user.Email, &hash, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return User{}, UserError{Code: 104, Msg: "User don't exist"}
	}

	if !checkPasswordHash(user.Password, hash) {
		return User{}, UserError{Code: 105, Msg: "Incorrect password"}
	}
	user.Password = ""
	return user, err
}

// Get gets user from database by email
func (m Model) Register(user User) (User, error) {
	if !isValidPass(user.Password) {
		return User{}, UserError{Code: 100, Msg: "Invalid password"}
	}
	if !validEmail.MatchString(user.Email) {
		return User{}, UserError{Code: 101, Msg: "Invalid email"}
	}
	hash, err := hashPassword(user.Password)
	if err != nil {
		return User{}, err
	}
	sqlStatement := `INSERT INTO users (email, password) VALUES($1, $2) 
		RETURNING id, created_at`
	err = m.db.QueryRow(sqlStatement, user.Email, hash).Scan(
		&user.ID, &user.CreatedAt)
	if err, ok := err.(*pq.Error); ok {
		if err.Code == "23505" {
			return User{}, UserError{Code: 103, Msg: "User already exists"}
		}
	}
	user.Password = ""
	return user, err
}
