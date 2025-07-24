package dao

import (
	"context"

	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
)

//go:generate mockgen -source=read.go -destination=mock/read.go -package=mock
type Reader interface {
	Select(ctx context.Context, sqlStr string, args ...any) ([]db.Event, error)
	Count(ctx context.Context, sqlStr string, args ...any) (int64, error)
}
