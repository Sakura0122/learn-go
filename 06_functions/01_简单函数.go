package main

import "fmt"

func sayHello() {
	fmt.Println("Hello, World!")
}

func add(numberList ...int) int {
	sum := 0
	for _, number := range numberList {
		sum += number
	}
	return sum
}

// 多个返回值
func getTwoNumbers() (int, int, int) {
	return 5, 10, 15
}

func main() {
	sayHello()
	fmt.Println(add(5, 3))
	fmt.Println(getTwoNumbers())
}
