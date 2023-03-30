package main

import "sync/atomic"

type Metrics struct {
	RequestCount int64
	ErrorCount   int64
}

func (m *Metrics) IncrementRequestCount() {
	atomic.AddInt64(&m.RequestCount, 1)
}

func (m *Metrics) IncrementErrorCount() {
	atomic.AddInt64(&m.ErrorCount, 1)
}

func (m *Metrics) GetMetrics() (int64, int64) {
	return atomic.LoadInt64(&m.RequestCount), atomic.LoadInt64(&m.ErrorCount)
}
