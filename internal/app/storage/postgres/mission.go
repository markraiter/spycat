package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/markraiter/spycat/internal/app/storage"
	"github.com/markraiter/spycat/internal/domain"
)

func (s *Storage) SaveMission(ctx context.Context, tx *sql.Tx, mission *domain.Mission) (int, error) {
	const op = "storage.SaveMission"

	query := "INSERT INTO missions (cat_id, notes, completed) VALUES ($1, $2, $3) RETURNING id"

	var missionID int
	err := tx.QueryRowContext(ctx, query, mission.CatID, mission.Notes, mission.Completed).Scan(&missionID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return missionID, nil
}

func (s *Storage) Missions(ctx context.Context) ([]*domain.Mission, error) {
	const op = "storage.Missions"

	query, err := s.PostgresDB.Prepare("SELECT id, cat_id, notes, completed FROM missions ORDER BY created_at DESC")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	missions := make([]*domain.Mission, 0)
	for rows.Next() {
		m := &domain.Mission{}
		err = rows.Scan(&m.ID, &m.CatID, &m.Notes, &m.Completed)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		missions = append(missions, m)
	}

	return missions, nil
}

func (s *Storage) MissionByID(ctx context.Context, id int) (*domain.Mission, error) {
	const op = "storage.MissionByID"

	query := "SELECT id, cat_id, notes, completed FROM missions WHERE id = $1"
	row := s.PostgresDB.QueryRowContext(ctx, query, id)

	m := &domain.Mission{}
	err := row.Scan(&m.ID, &m.CatID, &m.Notes, &m.Completed)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, storage.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return m, nil
}
