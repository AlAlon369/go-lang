package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	box := CodeWars.CreateBox(5, 8)
	for _, row := range box {
		fmt.Println(row)
	}
}
