package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "Total number of processed events",
	})

	randGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_rand_gauge",
		Help: "other thing",
	})
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func SetRandGauge() {
	go func() {
		for {
			randGauge.Set(rand.Float64())
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	recordMetrics()
	SetRandGauge()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
