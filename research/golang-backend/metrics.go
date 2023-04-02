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

func (m *Metrics) IncrementRequestCount(key string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if _, ok := m.requestCount[key]; !ok {
		m.requestCount[key] = 1
	} else {
		m.requestCount[key]++
	}
}

func (m *Metrics) IncrementErrorCount(key string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if _, ok := m.errorCount[key]; !ok {
		m.errorCount[key] = 1
	} else {
		m.errorCount[key]++
	}
}

func (m *Metrics) GetMetrics() map[string]interface{} {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return map[string]interface{}{
		"requestCount": m.requestCount,
		"errorCount":   m.errorCount,
	}
}
