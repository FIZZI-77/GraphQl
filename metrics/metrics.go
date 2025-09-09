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
			Buckets: prometheus.DefBuckets,
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
)

func RegisterMetrics() {
	prometheus.MustRegister(HTTPRequestsTotal)
	prometheus.MustRegister(HTTPRequestDuration)
	prometheus.MustRegister(GraphQLRequestsTotal)
	prometheus.MustRegister(GraphQLRequestDuration)
}
