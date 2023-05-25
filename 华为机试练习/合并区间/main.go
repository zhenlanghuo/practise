package main

import "sort"

func main() {}

type Interval struct {
	Start int
	End   int
}

/**
 *
 * @param intervals Interval类一维数组
 * @return Interval类一维数组
 */
func merge(intervals []*Interval) []*Interval {

	ans := make([]*Interval, 0)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start == intervals[j].Start {
			return intervals[i].End < intervals[j].End
		}
		return intervals[i].Start < intervals[j].Start
	})

	if len(intervals) == 0 {
		return ans
	}

	ans = append(ans, intervals[0])
	for _, interval := range intervals {
		back := ans[len(ans)-1]
		if interval.Start > back.End {
			ans = append(ans, interval)
		} else {
			back.End = max(back.End, interval.End)
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

