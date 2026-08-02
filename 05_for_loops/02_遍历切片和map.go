package main

import "fmt"

func main() {
	list1 := []string{"a", "b", "c", "d", "e"}
	for i := 0; i < len(list1); i++ {
		fmt.Println(list1[i])
	}

	for index, value := range list1 {
		fmt.Println("Index:", index, "Value:", value)
	}

	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range map1 {
		fmt.Println("Key:", key, "Value:", value)
	}
}
