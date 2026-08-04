package main

import (
	"encoding/json"
)

type Result[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := Person{Name: "Alice", Age: 25}
	result := Result[Person]{Code: 200, Data: p1, Message: "success"}
	res, _ := json.MarshalIndent(result, "", "  ")
	println(string(res))
}
