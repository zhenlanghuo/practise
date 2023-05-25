package main

import (
	"fmt"
	"sync"
)

var (
	print1Chan chan int
	printaChan chan int
	wg         sync.WaitGroup
)

func main() {

	print1Chan = make(chan int)
	printaChan = make(chan int)

	wg = sync.WaitGroup{}
	wg.Add(2)
	go print1()
	go printa()

	wg.Wait()
}

func print1() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		print1Chan <- 1
		<-printaChan
	}
}

func printa() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-print1Chan
		fmt.Println("a")
		printaChan <- 1
	}
}
