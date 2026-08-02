package main

import (
	"fmt"
	"slices"
)

func main() {
	nameList := []string{"jack", "lucky", "tom", "jerry", "mike"}

	fmt.Println(nameList)
	nameList = append(nameList, "mike")
	fmt.Println(nameList)

	list1 := make([]int, 5)
	fmt.Println(list1)

	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[:]
	slice2 := arr[1:3]
	fmt.Println(slice1, slice2)

	ints := []int{3, 2, 6, 4, 1}
	//sort.Ints(ints)
	slices.Sort(ints)
	fmt.Println(ints)
}
