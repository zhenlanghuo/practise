package ratelimit

import (
	"log"
	"sync"
	"time"
)

type LeakyBucketLimiter struct {
	//limit    int
	//unit     time.Duration
	interval time.Duration // 水流速速
	lastTime time.Time     // 上次请求时间
	mu       sync.Mutex
}

func NewLeakyBucketLimiter(limit int, unit time.Duration) *LeakyBucketLimiter {
	log.Printf("%v", unit/time.Duration(limit))
	return &LeakyBucketLimiter{
		interval: unit / time.Duration(limit),
		lastTime: time.Now(),
	}
}

func (l *LeakyBucketLimiter) TryAcquire() bool {

	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	//log.Printf("%v, %v", now.Sub(l.lastTime), l.interval)
	if now.Sub(l.lastTime) < l.interval {
		//log.Println("return false")
		return false
	}

	l.lastTime = now
	return true
}

type LeakyBucketLimiterWithSlack struct {
	interval      time.Duration // 水流速速
	lastTime      time.Time     // 上次请求时间
	lastRemaining time.Duration // 上次多余时间
	maxSlack      time.Duration // 最大多余时间
	mu            sync.Mutex
}

func NewLeakyBucketLimiterWithSlack(limit int, unit, maxSlack time.Duration) *LeakyBucketLimiterWithSlack {
	log.Printf("%v", unit/time.Duration(limit))
	return &LeakyBucketLimiterWithSlack{
		interval: unit / time.Duration(limit),
		lastTime: time.Now(),
		maxSlack: maxSlack,
	}
}

func (l *LeakyBucketLimiterWithSlack) TryAcquire() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	if now.Sub(l.lastTime)+l.lastRemaining < l.interval {
		return false
	}

	l.lastRemaining += now.Sub(l.lastTime) + l.lastRemaining - l.interval
	if l.lastRemaining > l.maxSlack {
		l.lastRemaining = l.maxSlack
	}

	l.lastTime = now
	return true
}
