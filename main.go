package main

import (
	"fmt"
	"sort"
)

func CanSortBalls(n int, containers [][]int) bool {
	TotalBallsByColor := make([]int, n)
	TotalBallsByContainer := make([]int, n)

	for containerNumber := range TotalBallsByContainer {
		for colorNumber := range TotalBallsByColor {
			TotalBallsByColor[colorNumber]		   += containers[containerNumber][colorNumber]
			TotalBallsByContainer[containerNumber] += containers[containerNumber][colorNumber]
		}
	}

// If we can construct a bijection between the number of balls of each color
// and the number of balls in each container, sorting is possible

	sort.Ints(TotalBallsByColor)
	sort.Ints(TotalBallsByContainer)

	for i := range TotalBallsByColor {
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
	for containerNumber := range containers {
		containers[containerNumber] = make([]int, n)
		for colorNumber := range containers[containerNumber] {
			fmt.Scan(&containers[containerNumber][colorNumber])
		}
	}

	if CanSortBalls(n, containers) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
