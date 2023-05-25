package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func RedisDistance(lonX, latX, lonY, latY float64) float64 {
	const radiansPerDegree = 0.01745329251994329576923690768488612713442871888541725456097191 // https://oeis.org/A019685
	const earthRadiusInMeters = 6372797.560856

	rLonX := lonX * radiansPerDegree
	rLatX := latX * radiansPerDegree
	rLonY := lonY * radiansPerDegree
	rLatY := latY * radiansPerDegree

	u := math.Sin((rLatY - rLatX) / 2.0)
	v := math.Sin((rLonY - rLonX) / 2.0)
	return 2.0 * earthRadiusInMeters *
		math.Asin(math.Sqrt(u*u+math.Cos(rLatX)*math.Cos(rLatY)*v*v))
}

// 美团-地理空间距离计算优化 https://tech.meituan.com/2014/09/05/lucene-distance.html
func DistanceSimplify(lonX, latX, lonY, latY float64) float64 {
	const radiansPerDegree = 0.01745329251994329576923690768488612713442871888541725456097191 // https://oeis.org/A019685
	const earthRadiusInMeters = 6372797.560856

	dLon := lonX - lonY // 经度差值
	dLat := latX - latY // 纬度差值

	aLat := (latX + latY) / 2 // 平均纬度

	ew := dLon * radiansPerDegree * earthRadiusInMeters * math.Cos(aLat*radiansPerDegree) // 东西距离
	sn := dLat * radiansPerDegree * earthRadiusInMeters                                   // 南北距离

	return math.Sqrt(ew*ew + sn*sn) // 用平面的矩形对角距离公式计算总距离
}

func main() {
	count := int64(0)
	wg := sync.WaitGroup{}

	runtime.GOMAXPROCS(4)

	quit := make(chan struct{})

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-quit:
					return
				default:
					for j := 0; j < 1000; j++ {
						DistanceSimplify(11.3, 11.2, 13.4, 13.5)
					}
					atomic.AddInt64(&count, 1)
				}
			}
		}()
	}

	time.Sleep(time.Second * 30)
	close(quit)
	wg.Wait()
	fmt.Println(count, float64(count)/float64(30))
}
