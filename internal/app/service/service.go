package service

import (
	"errors"
)

var (
	ErrAlreadyExists      = errors.New("already exists")
	ErrNotFound           = errors.New("not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrNotAllowed         = errors.New("user is not allowed to perform this operation")
)

type AuthStorage interface {
}

type Service struct {
}

func New(
	a AuthStorage,

) *Service {
	return &Service{}
}
