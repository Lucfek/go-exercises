package model

import (
	"database/sql"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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

type Token struct {
	UserID    uint64
	Email     string
	CreatedAt string
	*jwt.StandardClaims
}

// Login log in user
func (m Users) Login(user User) (string, error) {
	var hash string
	sqlStatement := `SELECT password, id, created_at FROM users WHERE email=$1`
	err := m.DB.QueryRow(sqlStatement, user.Email).Scan(&hash, &user.ID, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return "", ErrUserNotFound
	}

	if err = checkPasswordHash(user.Password, hash); err != nil {
		return "", ErrIncorrectPass
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	tk := &Token{
		UserID:    user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, err := token.SignedString([]byte("secret"))

	return tokenString, err
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
	sqlStatement := `INSERT INTO users (email, password) VALUES($1, $2) RETURNING id, created_at;`
	_, err = m.DB.Exec(sqlStatement, user.Email, hash)
	if err, ok := err.(*pq.Error); ok {
		if err.Code == "23505" {
			return ErrUserAlreadyExist
		}
	}

	return err
}
