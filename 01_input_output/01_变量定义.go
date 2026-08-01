package main

import (
	"fmt"
	"learn-go/01_input_output/version"
)

func main() {
	var name1 string = "jack"

	var name2 = "lucky"

	name3 := "ming"

	name4, name5 := "test1", "test2"

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(name4)
	fmt.Println(name5)
	fmt.Println(version.Version)

}
