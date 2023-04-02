package main

import (
	"sync"
)

type Metrics struct {
	requestCount map[string]int
	errorCount   map[string]int
	mutex        sync.Mutex
}

func NewMetrics() *Metrics {
	return &Metrics{
		requestCount: make(map[string]int),
		errorCount:   make(map[string]int),
	}
}

func (m *Metrics) IncrementRequestCount(route string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.requestCount[route]++
}

func (m *Metrics) IncrementErrorCount(route string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.errorCount[route]++
}

func (m *Metrics) GetMetrics() map[string]interface{} {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return map[string]interface{}{
		"requestCount": m.requestCount,
		"errorCount":   m.errorCount,
	}
}
