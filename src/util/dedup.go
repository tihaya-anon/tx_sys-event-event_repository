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
type DedupResult string

const (
	DEDUP_RESULT_NEW   DedupResult = "new"
	DEDUP_RESULT_DUP   DedupResult = "duplicated"
	DEDUP_RESULT_ERROR DedupResult = "error"
)

func IsDup(ctx context.Context, q Reader, dedupKey string) (*db.Event, DedupResult, error) {
	event, err := q.ReadEventByDedupKey(ctx, pgtype.Text{String: dedupKey, Valid: true})
	if err != nil && errors.Is(err, pgx.ErrNoRows) { //expected ErrNoRows, means not duplicated
		return nil, DEDUP_RESULT_NEW, nil
	}
	if err != nil { // unexpected error, isDup with default value false
		return nil, DEDUP_RESULT_ERROR, err
	}
	// otherwise, found and duplicated
	return &event, DEDUP_RESULT_DUP, nil
}
