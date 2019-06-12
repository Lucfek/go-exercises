package model

import "errors"

var ErrUserNotFound = errors.New("User not found")
var ErrUserAlreadyExist = errors.New("User already exist")
var ErrIncorrectPass = errors.New("Incorrect password")
var ErrInvalidPass = errors.New("Invalid password")
var ErrInvalidEmail = errors.New("Invalid Email")
