package main

import "fmt"

func add[T int | float32 | float64](num1, num2 T) T {
	return num1 + num2
}

func main() {

	a := add(0.1, 0.2)
	fmt.Println(a)
}
