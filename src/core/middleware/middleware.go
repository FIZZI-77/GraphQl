package middleware

import (
	"GraphQL/metrics"
	"context"
	"github.com/99designs/gqlgen/graphql"
	"time"
)

func GraphQLMetricsMiddleware() func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	return func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		opCtx := graphql.GetOperationContext(ctx)
		opName := opCtx.OperationName
		start := time.Now()

		res, err = next(ctx)

		metrics.GraphQLRequestsTotal.WithLabelValues(opName).Inc()
		metrics.GraphQLRequestDuration.WithLabelValues(opName).Observe(time.Since(start).Seconds())

		return res, err
	}
}
