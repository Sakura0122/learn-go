package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}
	result := add(1, 2)
	fmt.Println(result)
}
