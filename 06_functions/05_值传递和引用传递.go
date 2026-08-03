package main

import "fmt"

func changeNum(num *int) {
	fmt.Println(num)
	*num += 1
}

func main() {
	num := 10
	fmt.Println(&num)
	changeNum(&num)
	fmt.Println(num)
}
