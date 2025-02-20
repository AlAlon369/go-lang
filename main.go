package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	tests := [][]int{
		{12, 4}, {12, 5}, {10, 2}, {100, 10}, {101, 10},
	}

	for _, t := range tests {
		fmt.Printf("Solve(%d, %d) = %v\n", t[0], t[1], CodeWars.SolveFunc(t[0], t[1]))
	}
}
