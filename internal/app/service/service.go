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
	ErrMissionCompleted   = errors.New("this mission completed")
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
	MissionProvider
	MissionProcessor
}

type TargetStorage interface {
	TargetSaver
	TargetProcessor
}

type Service struct {
	AuthService
	CatService
	MissionService
	TargetService
}

func New(
	a AuthStorage,
	c CatStorage,
	m MissionStorage,
	t TargetStorage,

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
			provider:  m,
			processor: m,
		},
		TargetService: TargetService{
			saver:     t,
			processor: t,
		},
	}
}
