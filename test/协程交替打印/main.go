package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

func main() {

	queue := list.New()
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)
	finish := make(chan struct{})

	wg := sync.WaitGroup{}

	wg.Add(3)

	producer := func() {
		defer wg.Done()
		//mu.Lock()
		//defer mu.Unlock()
		for i := 0; i < 100; i++ {
			mu.Lock()
			queue.PushBack(i)
			cond.Signal()
			//cond.Wait()
			mu.Unlock()
			time.Sleep(time.Millisecond * 5)
		}
		close(finish)
		cond.Signal()
	}

	type handleFunc func(queue *list.List)

	A := func(queue *list.List) {
		front := queue.Front()
		v := front.Value.(int)
		if v%2 == 0 {
			fmt.Printf("A %v\n", v)
			queue.Remove(front)
		}
	}

	B := func(queue *list.List) {
		front := queue.Front()
		v := front.Value.(int)
		if v%2 == 1 {
			fmt.Printf("B %v\n", v)
			queue.Remove(front)
		}
	}

	consumer := func(hf handleFunc) {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		for {
			if queue.Len() == 0 {
				select {
				case <-finish:
					cond.Signal()
					return
				default:
					cond.Signal()
					cond.Wait()
				}
				continue
			}

			hf(queue)
			cond.Signal()
			cond.Wait()
		}
	}

	go producer()
	go consumer(A)
	go consumer(B)

	wg.Wait()

	fmt.Println("finish")
}
