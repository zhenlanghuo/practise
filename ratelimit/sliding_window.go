package ratelimit

import (
	"container/list"
	"errors"
	"sync"
	"time"
)

type SlidingWindowLimiter struct {
	mu                  sync.Mutex
	limit               int           // 窗口请求限制
	window              time.Duration // 窗口时间大小
	smallWindow         time.Duration // 小窗口时间大小
	smallWindows        int           // 小窗口个数
	smallWindowCounters *list.List    // 小窗口请求数量计数
	counter             int           // 当前窗口请求数量计数
}

type SmallWindowCounter struct {
	counter     int
	windowStart time.Time
}

func NewSlidingWindowLimiter(limit int, window, smallWindow time.Duration) (*SlidingWindowLimiter, error) {
	if window%smallWindow != 0 {
		return nil, errors.New("window cannot be split by integers")
	}

	return &SlidingWindowLimiter{
			limit:               limit,
			window:              window,
			smallWindow:         smallWindow,
			smallWindows:        int(window / smallWindow),
			smallWindowCounters: list.New(),
			mu:                  sync.Mutex{}},
		nil
}

func (l *SlidingWindowLimiter) TryAcquire(i int) bool {
	//log.Println(i, l.counter, l.limit)
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	for l.smallWindowCounters.Len() > 0 {
		top := l.smallWindowCounters.Front()
		topSmallWindowCounter := top.Value.(*SmallWindowCounter)
		if now.Sub(topSmallWindowCounter.windowStart) > l.window {
			l.smallWindowCounters.Remove(top)
			l.counter = l.counter - topSmallWindowCounter.counter
		} else {
			break
		}
	}

	if l.counter >= l.limit {
		//log.Println("false", i, l.counter, l.limit)
		return false
	}

	back := l.smallWindowCounters.Back()
	if back == nil {
		back = l.smallWindowCounters.PushBack(&SmallWindowCounter{windowStart: now, counter: 0})
	}
	backSmallWindowCounter := back.Value.(*SmallWindowCounter)
	if now.Sub(backSmallWindowCounter.windowStart) > l.smallWindow {
		interval := now.Sub(backSmallWindowCounter.windowStart) / l.smallWindow
		back = l.smallWindowCounters.PushBack(&SmallWindowCounter{windowStart: backSmallWindowCounter.windowStart.Add(interval * l.smallWindow), counter: 0})
	}
	back.Value.(*SmallWindowCounter).counter++
	l.counter++
	//log.Println(i, l.counter, l.limit, l.smallWindowCounters)
	//top := l.smallWindowCounters.Front()
	//for top != nil {
	//	log.Println(top.Value.(*SmallWindowCounter))
	//	top = top.Next()
	//}

	return true
}
