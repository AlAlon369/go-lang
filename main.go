package main

import (
	"Golang/CodeWars" // Импортируем пакет CodeWars
	"fmt"
)

func main() {
	// Тестирование метода IsUpperCase
	fmt.Println(CodeWars.MyString("c").IsUpperCase())                      // false
	fmt.Println(CodeWars.MyString("C").IsUpperCase())                      // true
	fmt.Println(CodeWars.MyString("hello I AM DONALD").IsUpperCase())      // false
	fmt.Println(CodeWars.MyString("HELLO I AM DONALD").IsUpperCase())      // true
	fmt.Println(CodeWars.MyString("ACSKLDFJSgSKLDFJSKLDFJ").IsUpperCase()) // false
	fmt.Println(CodeWars.MyString("ACSKLDFJSGSKLDFJSKLDFJ").IsUpperCase()) // true
}
