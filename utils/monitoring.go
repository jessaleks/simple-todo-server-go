package utils

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetUpMonitoring() {
	counter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "mtsouk",
			Name:      "my_counter",
			Help:      "This is my counter",
		})
	var gauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mtsouk",
			Name:      "my_gauge",
			Help:      "This is my gauge",
		})
	var histogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "mtsouk",
			Name:      "my_histogram",
			Help:      "This is my histogram",
		})
	var summary = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace: "mtsouk",
			Name:      "my_summary",
			Help:      "This is my summary",
		})
	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)
	go func() {
		for {
			counter.Add(rand.Float64() * 5)
			gauge.Add(rand.Float64()*15 - 5)
			histogram.Observe(rand.Float64() * 10)
			summary.Observe(rand.Float64() * 10)
			time.Sleep(2 * time.Second)
		}
	}()
}

func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}

}
