package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/markraiter/spycat/internal/app/storage"
	"github.com/markraiter/spycat/internal/domain"
	"github.com/markraiter/spycat/internal/lib/breed"
)

type CatSaver interface {
	SaveCat(ctx context.Context, cat *domain.Cat) (int, error)
}

type CatProvider interface {
	Cat(ctx context.Context, id int) (*domain.Cat, error)
	Cats(ctx context.Context) ([]*domain.Cat, error)
}

type CatProcessor interface {
	UpdateCat(ctx context.Context, cat *domain.Cat) error
	DeleteCat(ctx context.Context, id int) error
}

type CatService struct {
	saver     CatSaver
	provider  CatProvider
	processor CatProcessor
}

func (s *CatService) SaveCat(ctx context.Context, cr *domain.CatRequest) (int, error) {
	const op = "service.SaveCat"

	cat := &domain.Cat{
		Name:              cr.Name,
		YearsOfExperience: cr.YearsOfExperience,
		Breed:             cr.Breed,
		Salary:            cr.Salary,
	}

	if !breed.ValidateCatBreed(cat.Breed) {
		return 0, fmt.Errorf("%s: %w", op, ErrCatBreedNotFound)
	}

	id, err := s.saver.SaveCat(ctx, cat)
	if err != nil {
		if errors.Is(err, storage.ErrAlreadyExists) {
			return 0, fmt.Errorf("%s: %w", op, ErrAlreadyExists)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *CatService) Cat(ctx context.Context, id int) (*domain.Cat, error) {
	const op = "service.Cat"

	cat, err := s.provider.Cat(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return cat, nil
}

func (s *CatService) Cats(ctx context.Context) ([]*domain.Cat, error) {
	const op = "service.Cats"

	cats, err := s.provider.Cats(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return cats, nil
}

func (s *CatService) UpdateCat(ctx context.Context, catID int, cr *domain.CatRequest) error {
	const op = "service.UpdateCat"

	cat := &domain.Cat{
		ID:                catID,
		Name:              cr.Name,
		YearsOfExperience: cr.YearsOfExperience,
		Breed:             cr.Breed,
		Salary:            cr.Salary,
	}

	if !breed.ValidateCatBreed(cat.Breed) {
		return fmt.Errorf("%s: %w", op, ErrCatBreedNotFound)
	}

	if err := s.processor.UpdateCat(ctx, cat); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *CatService) DeleteCat(ctx context.Context, id int) error {
	const op = "service.DeleteCat"

	if err := s.processor.DeleteCat(ctx, id); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
