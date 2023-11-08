package main

import (
	"fmt"
	"github.com/zhshch2002/goreq"
	"github.com/zhshch2002/gospider"
	"time"
)

func main() {
	s := gospider.NewSpider() // create spider
	start := time.Now()
	s.AddRootTask(goreq.Get("https://adpemploymentreport.com/"), func(t *gospider.Task) {
		h, _ := t.HTML()
		fmt.Println(h.Find(".container-fluid.report-wrap.no-gutter .report-overview .stat .value").Text())
		fmt.Println(time.Since(start))
	})

	s.Wait()
}
