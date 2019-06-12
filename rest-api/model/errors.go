package model

import "fmt"

type UserError struct {
	Code int
	Msg  string
}

func (e UserError) Error() string {
	return fmt.Sprintf("%s, code: %d", e.Msg, e.Code)
}
