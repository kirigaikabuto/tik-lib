package tik_lib

import (
	"errors"
	com "github.com/kirigaikabuto/setdata-common"
)

var (
	ErrCreateUserUnknown      = com.NewMiddleError(errors.New("could not create user: unknown error"), 500, 1)
	ErrUserNotFound           = com.NewMiddleError(errors.New("user not found"), 404, 2)
	ErrUserIdNotProvided      = com.NewMiddleError(errors.New("user id is not provided"), 400, 3)
	ErrUserPasswordNotCorrect = com.NewMiddleError(errors.New("user password not correct"), 400, 4)
	ErrNothingToUpdate        = com.NewMiddleError(errors.New("nothing to update"), 400, 5)
)
