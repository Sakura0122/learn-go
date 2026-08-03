package main

import (
	"fmt"
	"time"
)

func awaitAdd(second int) func(...int) int {
	time.Sleep(time.Duration(second) * time.Second)

	return func(numbers ...int) int {
		sum := 0
		for _, number := range numbers {
			sum += number
		}
		return sum
	}
}

func main() {
	fmt.Println(awaitAdd(3)(1, 2, 3))
}
