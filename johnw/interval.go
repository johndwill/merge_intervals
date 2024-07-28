package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 2}, {3, 8}, {4, 5}, {1, 7}}
	// intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

	// intervals := [][]int{{1, 4}, {4, 5}}
	//intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

	output := merge(intervals)
	fmt.Println(output)
}

func merge(intervals [][]int) [][]int {
	// fmt.Println(intervals)
	if len(intervals) == 0 {
		return intervals
	}

	// sort.Slice(intervals, func(i, j int) bool {
	// 	return intervals[i][1] < intervals[j][1]
	// })
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// fmt.Println(intervals)

	output := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		curInterval := output[len(output)-1]
		thisInterval := intervals[i]

		if thisInterval[0] < curInterval[0] && thisInterval[1] >= curInterval[0] {
			curInterval[0] = thisInterval[0]
			continue
		} else if thisInterval[0] <= curInterval[1] && thisInterval[1] > curInterval[0] {
			if thisInterval[1] > curInterval[1] {
				curInterval[1] = thisInterval[1]
				continue
			}
		} else if thisInterval[0] > curInterval[1] {
			output = append(output, thisInterval)
		}
	}

	return output
}
