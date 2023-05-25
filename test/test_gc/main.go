package main

import (
	"arena"
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type Data struct {
	i1        int
	i2        int
	i3        int
	i4        int
	i5        int
	i6        int
	i7        int
	i8        int
	i9        int
	i10       int
	s1        []byte
	s2        []byte
	s3        []byte
	s4        []byte
	s5        []byte
	s6        []byte
	s7        []byte
	s8        []byte
	s9        []byte
	s10       []byte
	byteArray []byte
}

func NewData(a *arena.Arena, moreByte bool) *Data {
	var d *Data
	if a == nil {
		d = &Data{}
	} else {
		d = arena.New[Data](a)
	}

	d.i1 = 1
	d.i2 = 2
	d.i3 = 3
	d.i4 = 4
	d.i5 = 5
	d.i6 = 6
	d.i7 = 7
	d.i8 = 8
	d.i9 = 9
	d.i10 = 10

	if moreByte {
		//byteArray := make([]byte, 1)
		//for i := 0; i < len(byteArray); i++ {
		//	byteArray[i] = '1'
		//}
		size := 400

		d.s1 = make([]byte, size)
		d.s2 = make([]byte, size)
		d.s3 = make([]byte, size)
		d.s4 = make([]byte, size)
		d.s5 = make([]byte, size)
		d.s6 = make([]byte, size)
		d.s7 = make([]byte, size)
		d.s8 = make([]byte, size)
		d.s9 = make([]byte, size)
		d.s10 = make([]byte, size)
	} else {
		//byteArray := make([]byte, 1)
		//for i := 0; i < len(byteArray); i++ {
		//	byteArray[i] = '1'
		//}

		size := 10

		d.s1 = make([]byte, size)
		d.s2 = make([]byte, size)
		d.s3 = make([]byte, size)
		d.s4 = make([]byte, size)
		d.s5 = make([]byte, size)
		d.s6 = make([]byte, size)
		d.s7 = make([]byte, size)
		d.s8 = make([]byte, size)
		d.s9 = make([]byte, size)
		d.s10 = make([]byte, size)
	}

	//d.byteArray = make([]byte, 4000)
	return d
}

func main() {

	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	//a := arena.NewArena()
	//defer runtime.KeepAlive(a)

	size := int(1e7)
	maxLifeTime := time.Minute * 2

	m := storeWithMap(size)
	//m := storeWithMapWithArena(size, a)
	defer runtime.KeepAlive(m)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s\n", time.Since(start))
	}

	quit := make(chan bool)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	wg := sync.WaitGroup{}

	//sizes := []int{
	//	1024 * 5, 1024 * 5, 1024 * 5, 1024 * 5, 1024 * 5, 1024 * 5, 1024 * 5, 1024 * 5, 1024 * 5, 1024 * 5,
	//	1024 * 50, 1024 * 50, 1024 * 50, 1024 * 50, 1024 * 50, 1024 * 50, 1024 * 50, 1024 * 50, 1024 * 50, 1024 * 50,
	//	1024 * 500, 1024 * 500, 1024 * 500, 1024 * 500, 1024 * 500, 1024 * 500, 1024 * 500, 1024 * 500, 1024 * 500, 1024 * 500,
	//	1024 * 1024, 1024 * 1024, 1024 * 1024, 1024 * 1024, 1024 * 1024,
	//	1024 * 1024 * 10, 1024 * 1024 * 10, 1024 * 1024 * 10, 1024 * 1024 * 10, 1024 * 1024 * 10,
	//	1024 * 1024 * 50, 1024 * 1024 * 50,
	//	1024 * 1024 * 100,
	//}

	//applyWorkers := len(sizes)
	applyWorkers := 100

	costStaMap := make(map[int]*int64)
	for i := 0; i <= 10; i++ {
		costStaMap[i] = new(int64)
	}
	for i := 0; i <= 90; i += 10 {
		costStaMap[i] = new(int64)
	}
	for i := 100; i <= 1000; i += 100 {
		costStaMap[i] = new(int64)
	}

	wg.Add(applyWorkers)
	//for i := 0; i < len(sizes); i++ {
	//	go func(i, size int) {
	//		defer wg.Done()
	//		count := 0
	//		for {
	//			select {
	//			case <-quit:
	//				fmt.Printf("i: %v, count: %v\n", i, count)
	//				return
	//			default:
	//				_, cost := applyMemory(size)
	//				k := int(cost / time.Millisecond)
	//				if k < 10 {
	//					atomic.AddInt64(costStaMap[0], 1)
	//				} else if k >= 10 && k < 100 {
	//					atomic.AddInt64(costStaMap[k/10*10], 1)
	//				} else if k >= 100 && k < 1000 {
	//					atomic.AddInt64(costStaMap[k/100*100], 1)
	//				} else {
	//					atomic.AddInt64(costStaMap[1000], 1)
	//				}
	//
	//				time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	//				count++
	//			}
	//		}
	//	}(i, sizes[i])
	//}

	qps := 3000
	for i := 0; i < applyWorkers; i++ {
		go func(i, size, qps int) {
			defer wg.Done()
			count := 0
			limit := rate.NewLimiter(rate.Limit(qps), 1)
			for {
				select {
				case <-quit:
					fmt.Printf("i: %v, count: %v\n", i, count)
					return
				default:
					limit.Wait(context.Background())
					_, cost := applyMemory(size)
					k := int(cost / time.Millisecond)
					if k < 10 {
						atomic.AddInt64(costStaMap[k], 1)
					} else if k >= 10 && k < 100 {
						atomic.AddInt64(costStaMap[k/10*10], 1)
					} else if k >= 100 && k < 1000 {
						atomic.AddInt64(costStaMap[k/100*100], 1)
					} else {
						atomic.AddInt64(costStaMap[1000], 1)
					}
					//time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
					count++
				}
			}
		}(i, 200, qps/applyWorkers)
	}

	go func() {
		timer := time.NewTimer(maxLifeTime)
		select {
		case sig := <-sigs:
			fmt.Println()
			fmt.Println(sig)
		case <-timer.C:
			fmt.Println("到时间")
		}

		close(quit)
	}()

	fmt.Println("awaiting signal")
	<-quit
	wg.Wait()
	fmt.Println("exiting")

	for i := 0; i <= 9; i++ {
		fmt.Printf("%vms: %v\n", i, *costStaMap[i])
	}
	for i := 10; i <= 90; i += 10 {
		fmt.Printf("%vms: %v\n", i, *costStaMap[i])
	}
	for i := 100; i <= 1000; i += 100 {
		fmt.Printf("%vms: %v\n", i, *costStaMap[i])
	}
}

//func applyMemory(size int) ([]byte, time.Duration) {
//	start := time.Now()
//	bytes := make([]byte, size)
//	if time.Since(start) > time.Millisecond*20 {
//		fmt.Printf("applyMemory, size: %v, cost: %v\n", size, time.Since(start))
//	}
//
//	return bytes, time.Since(start)
//}

func applyMemory(size int) ([]*Data, time.Duration) {
	start := time.Now()

	randNum := rand.Intn(100)

	if randNum < 20 {
		size = size * 3
	} else if randNum < 50 {
		size = size * 2
	}

	bytes := make([]*Data, size)
	for i := 0; i < size; i++ {
		bytes[i] = NewData(nil, true)
	}
	if time.Since(start) > time.Millisecond*20 {
		fmt.Printf("applyMemory, size: %v, cost: %v\n", size, time.Since(start))
	}

	return bytes, time.Since(start)
}

func storeWithMap(size int) map[int]*Data {
	m := make(map[int]*Data, size)
	for i := 0; i < size; i++ {
		m[i] = NewData(nil, false)
	}
	return m
}

func storeWithMapWithArena(size int, a *arena.Arena) map[int]*Data {
	m := make(map[int]*Data, size)
	for i := 0; i < size; i++ {
		m[i] = NewData(a, false)
	}
	return m
}
