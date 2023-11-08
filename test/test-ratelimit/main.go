package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	limiter := rate.NewLimiter(100, 1)

	//limiter := ratelimit.NewTokenBucketLimiter(100, time.Second, 100)
	quit := make(chan struct{})
	wg := sync.WaitGroup{}
	count := int64(0)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-quit:
					return
				default:
				}
				if limiter.Allow() {
					atomic.AddInt64(&count, 1)
				}
				//limiter.Wait(context.Background())
				//atomic.AddInt64(&count, 1)
			}
		}()
	}

	time.Sleep(time.Second * 10)
	close(quit)
	wg.Wait()
	fmt.Println(count / 10)
}
