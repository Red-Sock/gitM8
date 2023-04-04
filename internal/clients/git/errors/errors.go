package errors

import (
	"github.com/pkg/errors"
)

var (
	ErrCouldNotFindCurrentUser = errors.New("couldn't find current user")
	ErrUnauthorized            = errors.New("invalid token: user unauthorized")
	ErrInvalidResponseData     = errors.New("git system responded with invalid data")
)
