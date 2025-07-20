package dao

import (
	"context"

	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
)

func Read(ctx context.Context, tx db.DBTX, sqlStr string, args ...any) ([]db.Event, error) {
	rows, err := tx.Query(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []db.Event
	for rows.Next() {
		var i db.Event
		if err := rows.Scan(
			&i.EventID,
			&i.EventTopic,
			&i.EventType,
			&i.Source,
			&i.CreatedAt,
			&i.ExpiresAt,
			&i.Status,
			&i.RetryCount,
			&i.DedupKey,
			&i.Metadata,
			&i.Payload,
			&i.TargetService,
			&i.CorrelationID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
func Count(ctx context.Context, tx db.DBTX, sqlStr string, args ...any) (int64, error) {
	row := tx.QueryRow(ctx, sqlStr, args...)
	var count int64
	err := row.Scan(&count)
	return count, err
}
