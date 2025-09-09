package middlewareMs

import (
	"GraphQL/metrics"
	"context"
	"github.com/99designs/gqlgen/graphql"
	"time"
)

func GraphQLFieldMetrics(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	start := time.Now()

	res, err = next(ctx)

	fieldCtx := graphql.GetFieldContext(ctx)
	fieldName := fieldCtx.Field.Name
	if fieldName == "" {
		fieldName = "unknown_field"
	}

	metrics.GraphQLRequestsTotal.WithLabelValues(fieldName).Inc()
	metrics.GraphQLRequestDuration.WithLabelValues(fieldName).
		Observe(time.Since(start).Seconds())

	return res, err
}

func GraphQLResponseMetrics(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	start := time.Now()

	resp := next(ctx)

	opCtx := graphql.GetOperationContext(ctx)
	opName := opCtx.OperationName
	if opName == "" {
		opName = "unknown_operation"
	}

	metrics.GraphQLRequestsTotal.WithLabelValues(opName).Inc()
	metrics.GraphQLRequestDuration.WithLabelValues(opName).
		Observe(time.Since(start).Seconds())

	return resp
}
