package util

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
)

type PageQuery struct {
	PageSql   string
	PageArgs  []any
	TotalSql  string
	TotalArgs []any
}

type DefaultValues struct {
	OrderBy string
	Limit   uint64
}

func BuildQueryFromProto(statement sq.StatementBuilderType, tableName string, dv *DefaultValues, q *pb.Query) (*PageQuery, error) {
	// Dynamic WHERE clause
	statement = addFilter(statement, q.Filters)

	// Total SQL + args
	totalSql, totalArgs, err := statement.Select("COUNT(*)").From(tableName).ToSql()
	if err != nil {
		return nil, err
	}

	builder := statement.Select(q.SelectFields...).From(tableName)
	// Dynamic ORDER BY
	builder = addOrderBy(builder, dv.OrderBy, q.OrderBy)

	// LIMIT
	builder = addLimit(builder, dv.Limit, q.Limit)

	// page token
	builder = decodePageToken(builder, q.PageToken)

	// Page SQL + args
	pageSqlStr, pageArgs, err := builder.ToSql()

	if err != nil {
		return nil, err
	}
	return &PageQuery{TotalSql: totalSql, TotalArgs: totalArgs, PageSql: pageSqlStr, PageArgs: pageArgs}, nil
}
func addFilter(statement sq.StatementBuilderType, filters []*pb.Query_Filter) sq.StatementBuilderType {
	for _, f := range filters {
		statement = convertFilter(statement, f)
	}
	return statement
}
func addLimit(builder sq.SelectBuilder, defaultLimit uint64, limit int32) sq.SelectBuilder {
	if limit <= 0 {
		return builder.Limit(defaultLimit)
	}
	return builder.Limit(uint64(limit))
}
func addOrderBy(builder sq.SelectBuilder, defaultOrderBy string, orderBy []*pb.Query_OrderBy) sq.SelectBuilder {
	if len(orderBy) == 0 {
		return builder.OrderBy(defaultOrderBy)
	}
	var orderExprs []string
	for _, o := range orderBy {
		buildOrderExprs(&orderExprs, o)
	}
	return builder.OrderBy(orderExprs...)
}
func buildOrderExprs(orderExprs *[]string, orderBy *pb.Query_OrderBy) {
	if orderBy.Field == "" {
		return
	}
	order := orderBy.Field
	if orderBy.Descending {
		order += " DESC"
	}
	*orderExprs = append(*orderExprs, order)
}
func convertFilter(statement sq.StatementBuilderType, f *pb.Query_Filter) sq.StatementBuilderType {
	field := f.Field
	if field == "" || len(f.Values) == 0 {
		return statement
	}
	switch f.Op {
	case pb.Query_Filter_EQ:
		return statement.Where(sq.Eq{field: f.Values[0]})
	case pb.Query_Filter_NEQ:
		return statement.Where(sq.NotEq{field: f.Values[0]})
	case pb.Query_Filter_LT:
		return statement.Where(sq.Lt{field: f.Values[0]})
	case pb.Query_Filter_LTE:
		return statement.Where(sq.LtOrEq{field: f.Values[0]})
	case pb.Query_Filter_GT:
		return statement.Where(sq.Gt{field: f.Values[0]})
	case pb.Query_Filter_GTE:
		return statement.Where(sq.GtOrEq{field: f.Values[0]})
	case pb.Query_Filter_IN:
		return statement.Where(sq.Eq{field: f.Values})
	case pb.Query_Filter_LIKE:
		return statement.Where(sq.Like{field: f.Values[0]})
	}
	return statement
}

// TODO implement page token
func decodePageToken(builder sq.SelectBuilder, token string) sq.SelectBuilder {
	if token == "" {
		return builder
	}
	var offset int
	if _, err := fmt.Sscanf(token, "page_%d", &offset); err != nil {
		return builder
	}
	builder = builder.Offset(uint64(offset))
	return builder
}

// TODO implement page token
func EncodePageToken(dv *DefaultValues, q *pb.Query) string {
	return fmt.Sprintf("page_%s", q.PageToken)
}
