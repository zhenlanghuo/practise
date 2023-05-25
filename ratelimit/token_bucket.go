package ratelimit

import (
	"sync"
	"time"
)

type TokenBucketLimiter struct {
	limit    int
	unit     time.Duration
	interval time.Duration // 每 interval 时间产生一个token
	lastTime time.Time     // 上次请求时间
	tokens   int           // 剩余token
	size     int           // 桶大小
	mu       sync.Mutex
}

func NewTokenBucketLimiter(limit int, unit time.Duration, size int) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		limit:    limit,
		unit:     unit,
		interval: unit / time.Duration(limit),
		lastTime: time.Now(),
		size:     size,
	}
}

func (l *TokenBucketLimiter) TryAcquire() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	newTokens := int(now.Sub(l.lastTime) / l.interval)
	l.tokens += newTokens
	//log.Println("TryAcquire", l.tokens, l.size, l.lastTime)
	if l.tokens > l.size {
		l.tokens = l.size
	}

	if l.tokens == 0 {
		return false
	}
	l.lastTime = now
	l.tokens--

	return true
}
