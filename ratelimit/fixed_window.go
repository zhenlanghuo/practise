package ratelimit

import (
	"sync"
	"time"
)

type FixedWindowLimiter struct {
	mu          sync.Mutex
	window      time.Duration // 窗口时间大小
	limit       int           // 窗口请求限制
	windowStart time.Time     // 窗口开始时间
	counter     int           // 窗口请求计数器
}

func NewFixedWindowLimiter(window time.Duration, limit int) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		window:      window,
		limit:       limit,
		windowStart: time.Now(),
	}
}

func (l *FixedWindowLimiter) TryAcquire() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	if now.Sub(l.windowStart) > l.window {
		l.windowStart = now
		l.counter = 0
	}

	if l.counter >= l.limit {
		return false
	}
	l.counter++
	return true
}
