package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Total HTTP requests counter
	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "route"},
	)

	// Latency Histogram
	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_duration_seconds",
			Help:    "Histogram of HTTP request durations",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "route"},
	)

	// Error Counter
	httpErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total number of HTTP request errors",
		},
		[]string{"method", "route"},
	)
)

func RecordMetrics(c *gin.Context) {
	// Track HTTP request duration
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		httpDuration.WithLabelValues(c.Request.Method, c.FullPath()).Observe(duration)

		// Increment error count if the response status code is >= 400
		if c.Writer.Status() >= 400 {
			httpErrors.WithLabelValues(c.Request.Method, c.FullPath()).Inc()
		}

		// Increment total request count
		httpRequests.WithLabelValues(c.Request.Method, c.FullPath()).Inc()
	}()
}

func init() {
	// Register the metrics
	prometheus.MustRegister(httpRequests)
	prometheus.MustRegister(httpDuration)
	prometheus.MustRegister(httpErrors)
}
