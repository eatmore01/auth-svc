package storage

import "errors"

var (
	UserExist    = errors.New("user already exist")
	UserNotFound = errors.New("user not found")
)
