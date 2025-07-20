package test

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-playground/assert"
)

func Test_SelectBuilder(t *testing.T) {
	// count
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	selectBuilderA := psql.Select("count(*)").From("events").Where(sq.Eq{"id": 1})
	selectBuilderB := psql.Where(sq.Eq{"id": 1}).Select("count(*)").From("events")
	sqlStrA, argsA, _ := selectBuilderA.ToSql()
	sqlStrB, argsB, _ := selectBuilderB.ToSql()
	assert.Equal(t, sqlStrA, sqlStrB)
	assert.Equal(t, argsA, argsB)
}
