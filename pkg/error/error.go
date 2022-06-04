package error

import "errors"

var (
	ErrReqNotAuthenticated = errors.New("request not authenticated")
)
