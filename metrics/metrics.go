package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	GraphQLRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "graphql_requests_total",
			Help: "Total number of GraphQL requests",
		},
		[]string{"path", "method"},
	)

	GraphQLRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "graphql_request_duration_seconds",
			Help:    "Duration of GraphQL requests",
			Buckets: prometheus.ExponentialBuckets(0.001, 2, 15),
		},
		[]string{"path", "method"},
	)
	HTTPRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method"},
	)

	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)

	DBQueriesTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "db_queries_total",
			Help: "Total number of database queries",
		},
		[]string{"query", "status"}, // status: success, error
	)

	DBQueryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_query_duration_seconds",
			Help:    "Duration of database queries",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"query"},
	)
)

func RegisterMetrics() {

	prometheus.MustRegister(HTTPRequestsTotal)
	prometheus.MustRegister(HTTPRequestDuration)
	prometheus.MustRegister(GraphQLRequestsTotal)
	prometheus.MustRegister(GraphQLRequestDuration)
	prometheus.MustRegister(DBQueriesTotal)
	prometheus.MustRegister(DBQueryDuration)

}
