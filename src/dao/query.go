package dao

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
)

//go:generate mockgen -source=query.go -destination=mock/query.go -package=mock
type Query interface {
	CreateEvent(ctx context.Context, arg db.CreateEventParams) error
	ReadEventByDedupKey(ctx context.Context, dedupKey pgtype.Text) (db.Event, error)
	ReadEventByEventId(ctx context.Context, eventID string) (db.Event, error)
	UpdateEventStatus(ctx context.Context, arg db.UpdateEventStatusParams) error
}
