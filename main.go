package main

import (
	"net/http"

	"github.com/JonasVgt/nvme-exporter/collector"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	c := collector.NewCollector()

	prometheus.MustRegister(c)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8000", nil)
}