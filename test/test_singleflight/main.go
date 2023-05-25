package main

import (
	"golang.org/x/sync/singleflight"
)

func main() {
	//singleflight.Group{}
}

type MultiSingleFlight struct {
	sf singleflight.Group
}

//func (m *MultiSingleFlight) Do(keys []string, fn func(keys []string) map[string]*singleflight.Result) map[string]*singleflight.Result {
//	left := make([]string, 0)
//	chanMap := make(map[string]chan *singleflight.Result)
//	resultMap := make(map[string]*singleflight.Result)
//	for _, key := range keys {
//		val, _, _ := m.sf.Do(key, func() (interface{}, error) {
//			left = append(left, key)
//			return make(chan *singleflight.Result, 1), nil
//		})
//		chanMap[key] = val.(chan *singleflight.Result)
//	}
//
//	leftResultMap := fn(left)
//	for key, result := range leftResultMap {
//		chanMap[key] <- result
//	}
//
//	for _, key := range keys {
//		resultMap[key] = <-chanMap[key]
//	}
//
//	return resultMap
//}

func (m *MultiSingleFlight) Do(keys []string, fn func(keys []string) map[string]*singleflight.Result) map[string]*singleflight.Result {
	left := make([]string, 0)
	chanMap := make(map[string]chan *singleflight.Result)
	resultMap := make(map[string]*singleflight.Result)
	for _, key := range keys {
		val, _, _ := m.sf.Do(key, func() (interface{}, error) {
			left = append(left, key)
			return make(chan *singleflight.Result, 1), nil
		})
		chanMap[key] = val.(chan *singleflight.Result)
	}

	leftResultMap := fn(left)
	for key, result := range leftResultMap {
		chanMap[key] <- result
	}

	for _, key := range keys {
		resultMap[key] = <-chanMap[key]
	}

	return resultMap
}
