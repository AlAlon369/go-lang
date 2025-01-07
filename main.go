package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	testCases := []struct {
		ages []int
	}{
		{[]int{1, 2, 10, 8}},
		{[]int{1, 5, 87, 45, 8, 8}},
		{[]int{1, 3, 10, 0}},
		{[]int{10, 15}},
		{[]int{6, 5, 83, 5, 3, 18}},
	}

	for _, tc := range testCases {
		result := CodeWars.TwoOldestAges(tc.ages)
		fmt.Printf("TwoOldestAges(%v) = %v\n", tc.ages, result)
	}
}
