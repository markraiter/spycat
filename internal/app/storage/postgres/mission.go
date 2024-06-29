package postgres

import (
	"context"
	"database/sql"
	"fmt"

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
