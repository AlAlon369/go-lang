package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	testCases := []struct {
		st string
		k  int
	}{
		{"1234", 1},
		{"1234", 2},
		{"1234", 3},
		{"2020", 1},
		{"987654321", 2},
	}

	for _, tc := range testCases {
		result := CodeWars.Solve(tc.st, tc.k)
		fmt.Printf("Solve(%q, %d) = %d\n", tc.st, tc.k, result)
	}
}
