package service

import (
	"errors"
)

var (
	ErrAlreadyExists      = errors.New("already exists")
	ErrNotFound           = errors.New("not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrCatBreedNotFound   = errors.New("cat breed not found")
)

type AuthStorage interface {
	UserSaver
	UserProvider
}

type CatStorage interface {
	CatSaver
	CatProvider
	CatProcessor
}

type Service struct {
	AuthService
	CatService
}

func New(
	a AuthStorage,
	c CatStorage,

) *Service {
	return &Service{
		AuthService: AuthService{
			saver:    a,
			provider: a,
		},
		CatService: CatService{
			saver:     c,
			provider:  c,
			processor: c,
		},
	}
}
