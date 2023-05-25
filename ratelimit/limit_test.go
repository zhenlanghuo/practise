package ratelimit

import (
	"log"
	"sync/atomic"
	"testing"
	"time"
)

func TestSlidingWindowLimit_TryAcquire(t *testing.T) {
	limiter, err := NewSlidingWindowLimiter2(5, time.Second, time.Millisecond*100)
	if err != nil {
		t.Fatal(err)
	}

	count := int32(0)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				if limiter.TryAcquire() {
					atomic.AddInt32(&count, 1)
					log.Printf("aa %v, %v", i, count)
				}
			}
			//for {
			//	log.Printf("aa %v", i)
			//}
		}(i)
	}

	time.Sleep(time.Second * 10)
}

func TestFixedWindowLimit_TryAcquire(t *testing.T) {
	limiter := NewFixedWindowLimiter(time.Second, 10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				if limiter.TryAcquire() {
					log.Printf("aa %v", i)
				}
			}
			//for {
			//	log.Printf("aa %v", i)
			//}
		}(i)
	}

	time.Sleep(time.Second * 10)
}

func TestLeakyBucketLimiter_TryAcquire(t *testing.T) {
	limiter := NewLeakyBucketLimiter(10, time.Second)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				if limiter.TryAcquire() {
					log.Printf("aa %v", i)
				}
			}
			//for {
			//	log.Printf("aa %v", i)
			//}
		}(i)
	}

	time.Sleep(time.Second * 10)
}

func TestLeakyBucketLimiterWithSlack_TryAcquire(t *testing.T) {
	limiter := NewLeakyBucketLimiterWithSlack(10, time.Second, time.Second)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				if limiter.TryAcquire() {
					log.Printf("aa %v", i)
				}
			}
			//for {
			//	log.Printf("aa %v", i)
			//}
		}(i)
	}

	time.Sleep(time.Second * 10)
}

func TestTokenBucketLimiter_TryAcquire(t *testing.T) {
	limiter := NewTokenBucketLimiter(10, time.Second, 50)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				if limiter.TryAcquire() {
					log.Printf("aa %v", i)
				}
			}
			//for {
			//	log.Printf("aa %v", i)
			//}
		}(i)
	}

	time.Sleep(time.Second * 10)
}
