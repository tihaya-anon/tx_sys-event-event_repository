package util

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
)

func IsDedup(ctx context.Context, q db.Queries, eventId, dedupKey string) (bool, error) {
	event, err := q.ReadEventByEventId(ctx, eventId)
	if err == pgx.ErrNoRows { // not found, expected, means not dedup
		return false, nil
	}
	if err != nil { // other error
		return true, err
	}
	if event.DedupKey.String != dedupKey {
		// found, but dedup key is different, so not dedup
		return false, nil
	}
	// otherwise, found and dedup
	return true, nil
}
