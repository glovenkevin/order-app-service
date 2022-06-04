package error

import "errors"

var (
	ErrReqNotAuthenticated = errors.New("Request not authenticated")
	ErrUserAlreadyExists   = errors.New("User already exists")
	ErrEmailAlreadyExists  = errors.New("Email already exists")
)
