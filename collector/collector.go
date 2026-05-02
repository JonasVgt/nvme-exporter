package collector

import (
	"math/rand"
	"time"

	"github.com/JonasVgt/nvme-exporter/metrics"
	
	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct{}

func NewCollector() *Collector {
	metrics.Init()

	c := &Collector{}
	go c.loop()
	return c
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {

}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {}

func (c *Collector) loop() {
	for {
		metrics.RandomValue.Set(rand.Float64())
		time.Sleep(5 * time.Second)
	}
}