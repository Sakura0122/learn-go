package main

import "fmt"

func main() {
	//map1 := make(map[string]int)
	map1 := map[string]int{}
	map1["key1"] = 1
	map1["key2"] = 2
	fmt.Println(map1)
	fmt.Println(map1["key3"])

	value, exists := map1["key3"]
	fmt.Println(value, exists)
}
