package model

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyExist = errors.New("user already exist")
var ErrIncorrectPass = errors.New("incorrect password")
var ErrInvalidPass = errors.New("invalid password")
var ErrInvalidEmail = errors.New("invalid Email")
