package handler

import (
	"GraphQL/metrics"
	"net/http"
	"time"
)

func HTTPMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start).Seconds()
		path := r.URL.Path
		method := r.Method

		metrics.HTTPRequestsTotal.WithLabelValues(path, method).Inc()
		metrics.HTTPRequestDuration.WithLabelValues(path, method).Observe(duration)
	})
}
