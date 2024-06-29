package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/markraiter/spycat/internal/app/storage"
)

type TargetProcessor interface {
	TargetCompleted(ctx context.Context, id int) error
	AddTargetToMission(ctx context.Context, missionID, targetID int) error
}

type TargetService struct {
	processor TargetProcessor
}

func (s *TargetService) CompleteTarget(ctx context.Context, id int) error {
	const op = "service.TargetCompleted"

	if err := s.processor.TargetCompleted(ctx, id); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *TargetService) AddTargetToMission(ctx context.Context, missionID, targetID int) error {
	const op = "service.AddTargetToMission"

	if err := s.processor.AddTargetToMission(ctx, missionID, targetID); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		if errors.Is(err, storage.ErrMissionCompleted) {
			return fmt.Errorf("%s: %w", op, ErrMissionCompleted)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
