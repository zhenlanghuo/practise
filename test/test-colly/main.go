package main

import (
	"context"
	"fmt"
	"github.com/gocolly/colly"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	cost := int64(0)
	count := int64(0)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
				atomic.AddInt64(&cost, int64(craw()))
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	wg.Wait()

	fmt.Printf("finish, cost: %v", time.Duration(cost/count))
}

func craw() time.Duration {
	var start time.Time
	//c := colly.NewCollector(colly.MaxDepth(3), colly.MaxBodySize(1024*28), colly.UserAgent("11"), colly.Async(true))
	c := colly.NewCollector(colly.MaxDepth(3))
	//c.OnResponse(func(response *colly.Response) {
	//	//fmt.Println(string(response.Body))
	//	//fmt.Println("OnResponse", time.Since(start))
	//})
	c.OnHTML(".container-fluid.report-wrap.no-gutter .report-overview .stat .value", func(element *colly.HTMLElement) {
		fmt.Println("OnHTML", time.Since(start), fmt.Sprintf("%+v", element))
	})

	start = time.Now()

	err := c.Request("GET", "https://adpemploymentreport.com/", nil, colly.NewContext(), nil)
	if err != nil {
		fmt.Println(err)
	}

	c.Wait()

	return time.Since(start)
}
