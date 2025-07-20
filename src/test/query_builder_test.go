package test

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-playground/assert"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/util"
)

func Test_QueryBuilder(t *testing.T) {
	// defualt
	query := &pb.Query{
		SelectFields: []string{"id", "name"},
		Filters: []*pb.Query_Filter{
			{Field: "id", Op: pb.Query_Filter_EQ, Values: []string{"1"}},
			{Field: "age", Op: pb.Query_Filter_LT, Values: []string{"100"}},
			{Field: "name", Op: pb.Query_Filter_EQ, Values: []string{"abc"}},
		},
	}
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	dv := util.DefaultValues{OrderBy: "create_at DESC", Limit: 100}
	pageQuery, err := util.BuildQueryFromProto(psql, "events", dv, query)
	assert.Equal(t, err, nil)
	assert.Equal(t, pageQuery.PageSql, "SELECT id, name FROM events WHERE id = $1 AND age < $2 AND name = $3 ORDER BY create_at DESC LIMIT 100")
	assert.Equal(t, pageQuery.PageArgs, []any{"1", "100", "abc"})
	assert.Equal(t, pageQuery.TotalSql, "SELECT COUNT(*) FROM events WHERE id = $1 AND age < $2 AND name = $3")
	assert.Equal(t, pageQuery.TotalArgs, []any{"1", "100", "abc"})

	// not defualt
	query.OrderBy = []*pb.Query_OrderBy{
		{Field: "id", Descending: true},
		{Field: "age", Descending: false},
	}
	query.Limit = 10

	pageQuery, err = util.BuildQueryFromProto(psql, "events", dv, query)
	assert.Equal(t, err, nil)
	assert.Equal(t, pageQuery.PageSql, "SELECT id, name FROM events WHERE id = $1 AND age < $2 AND name = $3 ORDER BY id DESC, age LIMIT 10")
}
