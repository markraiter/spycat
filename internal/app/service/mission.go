package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/markraiter/spycat/internal/domain"
)

type MissionSaver interface {
	SaveMission(ctx context.Context, tx *sql.Tx, mission *domain.Mission) (int, error)
	SaveTarget(ctx context.Context, tx *sql.Tx, target *domain.Target) error
}

type MissionProcessor interface {
	BeginTx(ctx context.Context) (*sql.Tx, error)
}

type MissionService struct {
	saver     MissionSaver
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
