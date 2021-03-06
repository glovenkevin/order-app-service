package error

import "errors"

var (
	ErrRequestCanceled  = errors.New("Request canceled")
	ErrDeadlineExceeded = errors.New("Deadline exceeded")

	ErrReqNotAuthenticated = errors.New("Request not authenticated")
	ErrUserNotFound        = errors.New("User not found")
	ErrUserAlreadyExists   = errors.New("User already exists")
	ErrEmailAlreadyExists  = errors.New("Email already exists")
	ErrPasswordIncorrect   = errors.New("Password incorrect")
)
