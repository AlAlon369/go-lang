package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	ch := make(chan int)
	go CodeWars.BaumSweet(ch)

	for i := 0; i < 20; i++ {
		fmt.Print(<-ch, " ")
	}
}
