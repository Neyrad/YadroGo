package main

import (
	"fmt"
	"sort"
)

func CanSortBalls(n int, containers [][]int) bool {
	TotalBallsByColor := make([]int, n)
	TotalBallsByContainer := make([]int, n)

	for containerNumber := 0; containerNumber < n; containerNumber++ {
		for colorNumber := 0; colorNumber < n; colorNumber++ {
			TotalBallsByColor[colorNumber]		   += containers[containerNumber][colorNumber]
			TotalBallsByContainer[containerNumber] += containers[containerNumber][colorNumber]
		}
	}

	sort.Ints(TotalBallsByColor)
	sort.Ints(TotalBallsByContainer)

	for i := 0; i < n; i++ {
		if TotalBallsByColor[i] != TotalBallsByContainer[i] {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)

	containers := make([][]int, n)
	for containerNumber := 0; containerNumber < n; containerNumber++ {
		containers[containerNumber] = make([]int, n)
		for colorNumber := 0; colorNumber < n; colorNumber++ {
			fmt.Scan(&containers[containerNumber][colorNumber])
		}
	}

	if CanSortBalls(n, containers) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
