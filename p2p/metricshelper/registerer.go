package metricshelper

import (
	"github.com/prometheus/client_golang/prometheus"
)

func RegisterCollectors(reg prometheus.Registerer, collectors ...prometheus.Collector) {
	_ = "STUB: not implemented"
	return
}
