package ratelimit

import (
	"errors"
	"sync"
	"time"
)

// SlidingWindowLimiter2 滑动窗口限流器
type SlidingWindowLimiter2 struct {
	limit        int           // 窗口请求上限
	window       time.Duration // 窗口时间大小
	smallWindow  time.Duration // 小窗口时间大小
	smallWindows int64         // 小窗口数量
	counters     map[int64]int // 小窗口计数器
	mutex        sync.Mutex    // 避免并发问题
}

// NewSlidingWindowLimiter 创建滑动窗口限流器
func NewSlidingWindowLimiter2(limit int, window, smallWindow time.Duration) (*SlidingWindowLimiter2, error) {
	// 窗口时间必须能够被小窗口时间整除
	if window%smallWindow != 0 {
		return nil, errors.New("window cannot be split by integers")
	}

	return &SlidingWindowLimiter2{
		limit:        limit,
		window:       window,
		smallWindow:  smallWindow,
		smallWindows: int64(window / smallWindow),
		counters:     make(map[int64]int),
	}, nil
}

func (l *SlidingWindowLimiter2) TryAcquire() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	now := time.Now()
	// 获取当前小窗口值
	currentSmallWindow := now.UnixNano() / int64(l.smallWindow) * int64(l.smallWindow)
	// 获取起始小窗口值
	startSmallWindow := currentSmallWindow - int64(l.smallWindow)*(l.smallWindows-1)

	// 计算当前窗口的请求总数
	var count int
	for smallWindow, counter := range l.counters {
		if smallWindow < startSmallWindow {
			delete(l.counters, smallWindow)
		} else {
			count += counter
		}
	}

	// 若到达窗口请求上限，请求失败
	if count >= l.limit {
		return false
	}

	// 若没到窗口请求上限，当前小窗口计数器+1，请求成功
	l.counters[currentSmallWindow]++
	return true
}
