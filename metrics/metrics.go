package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	RandomValue = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "my_random_value",
		Help: "A random value from the collector",
	})
)

func Init() {
	prometheus.MustRegister(RandomValue)
}