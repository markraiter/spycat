package service

import (
	"errors"
)

var (
	ErrAlreadyExists      = errors.New("already exists")
	ErrNotFound           = errors.New("not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrCatBreedNotFound   = errors.New("cat breed not found")
	ErrTooManyTargets     = errors.New("too many targets")
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

type MissionStorage interface {
	MissionSaver
	MissionProcessor
}

type Service struct {
	AuthService
	CatService
	MissionService
}

func New(
	a AuthStorage,
	c CatStorage,
	m MissionStorage,

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
		MissionService: MissionService{
			saver:     m,
			processor: m,
		},
	}
}
