package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/markraiter/spycat/internal/app/storage"
	"github.com/markraiter/spycat/internal/domain"
)

type MissionSaver interface {
	SaveMission(ctx context.Context, tx *sql.Tx, mission *domain.Mission) (int, error)
	SaveTarget(ctx context.Context, tx *sql.Tx, target *domain.Target) error
}

type MissionProvider interface {
	Missions(ctx context.Context) ([]*domain.Mission, error)
	MissionByID(ctx context.Context, id int) (*domain.Mission, error)
	Targets(ctx context.Context) ([]*domain.Target, error)
}

type MissionProcessor interface {
	BeginTx(ctx context.Context) (*sql.Tx, error)
}

type MissionService struct {
	saver     MissionSaver
	provider  MissionProvider
	processor MissionProcessor
}

func (s *MissionService) SaveMission(ctx context.Context, mr *domain.MissionRequest) (int, error) {
	const op = "service.SaveMission"

	mission := &domain.Mission{
		CatID:     mr.CatID,
		Targets:   mr.Targets,
		Notes:     mr.Notes,
		Completed: mr.Completed,
	}

	if len(mission.Targets) > 3 {
		return 0, fmt.Errorf("%s: %w", op, ErrTooManyTargets)
	}

	tx, err := s.processor.BeginTx(ctx)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	missionID, err := s.saver.SaveMission(ctx, tx, mission)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	for _, target := range mission.Targets {
		target.MissionID = missionID
		if err := s.saver.SaveTarget(ctx, tx, &target); err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("%s: %w", op, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return missionID, nil
}

func (s *MissionService) Missions(ctx context.Context) ([]*domain.Mission, error) {
	const op = "service.Missions"

	tx, err := s.processor.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	missions, err := s.provider.Missions(ctx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	targets, err := s.provider.Targets(ctx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for _, m := range missions {
		for _, t := range targets {
			if t.MissionID == m.ID {
				m.Targets = append(m.Targets, *t)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return missions, nil
}

func (s *MissionService) MissionByID(ctx context.Context, id int) (*domain.Mission, error) {
	const op = "service.MissionByID"

	tx, err := s.processor.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	mission, err := s.provider.MissionByID(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			tx.Rollback()
			return nil, fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		tx.Rollback()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	targets, err := s.provider.Targets(ctx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for _, t := range targets {
		if t.MissionID == mission.ID {
			mission.Targets = append(mission.Targets, *t)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return mission, nil
}
