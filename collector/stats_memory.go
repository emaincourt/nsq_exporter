package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

type memoryStats []struct {
	val func(*memory) float64
	vec *prometheus.GaugeVec
}

// MemoryStats creates a new stats collector which is able to
// expose the golang memory metrics of a nsqd node to Prometheus.
func MemoryStats(namespace string) StatsCollector {
	labels := []string{}
	namespace += "_memory"

	return memoryStats{
		{
			val: func(m *memory) float64 { return float64(m.HeapObject) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "heap_object",
				Help:      "Heap object",
			}, labels),
		},
		{
			val: func(m *memory) float64 { return float64(m.HeapIdleBytes) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "heap_idle_bytes",
				Help:      "Heap idle bytes",
			}, labels),
		},
		{
			val: func(m *memory) float64 { return float64(m.HeapInUseBytes) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "heap_in_use_bytes",
				Help:      "Heap in use bytes",
			}, labels),
		},
		{
			val: func(m *memory) float64 { return float64(m.HeapReleasedBytes) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "heap_released_bytes",
				Help:      "Heap released bytes",
			}, labels),
		},
		{
			val: func(m *memory) float64 { return float64(m.GCPauseUsec100) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "gc_pause_usec_100",
				Help:      "GC pause usec 100",
			}, labels),
		},
		{
			val: func(m *memory) float64 { return float64(m.GCPauseUsec99) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "gc_pause_usec_99",
				Help:      "GC pause usec 99",
			}, labels),
		},
		{
			val: func(m *memory) float64 { return float64(m.GCPauseUsec95) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "gc_pause_usec_95",
				Help:      "GC pause usec 95",
			}, labels),
		},
		{
			val: func(m *memory) float64 { return float64(m.NextGCBytes) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "next_gc_bytes",
				Help:      "Next GC bytes",
			}, labels),
		},
		{
			val: func(m *memory) float64 { return float64(m.GCTotalRuns) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "gc_total_runs",
				Help:      "GC total runs",
			}, labels),
		},
	}
}

func (ms memoryStats) set(s *stats) {
	for _, c := range ms {
		c.vec.With(prometheus.Labels{}).Set(c.val(s.Memory))
	}
}

func (ms memoryStats) collect(out chan<- prometheus.Metric) {
	for _, c := range ms {
		c.vec.Collect(out)
	}
}

func (ms memoryStats) describe(ch chan<- *prometheus.Desc) {
	for _, c := range ms {
		c.vec.Describe(ch)
	}
}

func (ms memoryStats) reset() {
	for _, c := range ms {
		c.vec.Reset()
	}
}
