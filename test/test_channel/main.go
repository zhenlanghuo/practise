package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int, 10)
	c <- 1
	close(c)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		//L:
		for {
			select {
			case i, _ := <-c:
				//if !ok {
				//	fmt.Println("break")
				//	break L
				//}
				fmt.Println(i)
			}
		}
		fmt.Println("!!")
	}()

	//time.Sleep(2)
	fmt.Println("&&&")
	wg.Wait()

	//context.WithCancel()
	//context.WithDeadline()
	//context.WithValue()
	sync.Map{}
}
