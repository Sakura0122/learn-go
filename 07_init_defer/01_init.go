package main

import "fmt"

var db string

func init() {
	db = "mysql"
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func init() {
	fmt.Println("init3")
}

func main() {
	fmt.Println("main")
	fmt.Println(db)
}
