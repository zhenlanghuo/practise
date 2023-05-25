package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var randMax = 2147483647

//var randMax = 32767

// var randMax = int64(math.MaxInt64)
var lfuLogFactor = 10

func main() {
	fmt.Println(math.MaxInt32, math.MaxInt64)
	rand.Seed(time.Now().Unix())
	result := make([]int, 255)
	for target := range result {
		for i := 0; i < 100; i++ {
			counter := uint8(0)
			count := 0
			for counter != uint8(target+1) {
				counter = LFULogIncr(counter)
				count++
			}
			result[target] += count
		}
		result[target] = result[target] / 100
	}
	for i, v := range result {
		fmt.Printf("%v: %v\n", i, v)
	}
}

func LFULogIncr(counter uint8) uint8 {
	if counter == 255 {
		return 255
	}

	//r := float64(rand.Int63n(randMax)) / float64(randMax)
	r := float64(rand.Intn(randMax)) / float64(randMax)
	baseval := counter - 5
	if baseval < 0 {
		baseval = 0
	}
	p := 1.0 / float64((int(baseval)*lfuLogFactor)+1)
	if r < p {
		counter++
	}
	return counter
}
