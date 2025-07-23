package util

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
)

type Reader interface {
	ReadEventByDedupKey(ctx context.Context, dedupKey pgtype.Text) (db.Event, error)
}

func IsDup(ctx context.Context, q Reader, dedupKey string) (*db.Event, error) {
	event, err := q.ReadEventByDedupKey(ctx, pgtype.Text{String: dedupKey, Valid: true})
	// error occurred
	if err != nil && errors.Is(err, pgx.ErrNoRows) { //expected ErrNoRows, means not duplicated
		return nil, nil
	}
	if err != nil { // unexpected error, isDup with default value false
		return nil, err
	}
	// otherwise, found and duplicated
	return &event, nil
}
