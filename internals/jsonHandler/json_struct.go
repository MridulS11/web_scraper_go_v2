package jsonhandler

import (
	"sync"
	"sync/atomic"
	"time"
)

type Metrics struct {
	Urls_Processed atomic.Uint64 `json:"urls_processed"`
	ErrorsEncountered atomic.Uint64 `json:"Errors"`
	BytesDownloaded int `json:"Memory"`
	TimeTaken time.Duration `json:"Time_Consumed"`
	mu sync.Mutex
}

func (m * Metrics) IncrementUrls() uint64{
	return m.Urls_Processed.Add(1)
}

func (m * Metrics) IncrementErrors() uint64 {
	return m.ErrorsEncountered.Add(1)
}

func (m * Metrics) IncrementBytes(bytes int) int{
	m.mu.Lock()
	defer m.mu.Unlock()
	m.BytesDownloaded += bytes
	return m.BytesDownloaded
}

func (m * Metrics) IncrementTime(t time.Duration) time.Duration{
	m.mu.Lock()
	defer m.mu.Unlock()
	m.TimeTaken = m.TimeTaken + t
	return m.TimeTaken
}

type MetricsSnapshot struct {
	UrlsProcessed     uint64  `json:"urls_processed"`
	ErrorsEncountered uint64  `json:"errors"`
	BytesDownloaded   int `json:"memory"`
	TimeTaken         time.Duration     `json:"time_consumed"`
}

func (m * Metrics) Snapshot() MetricsSnapshot {
	return MetricsSnapshot{
		UrlsProcessed:     m.Urls_Processed.Load(),
		ErrorsEncountered: m.ErrorsEncountered.Load(),
		BytesDownloaded:   m.BytesDownloaded,
		TimeTaken:         m.TimeTaken,
	}
}