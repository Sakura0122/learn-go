package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"-"`
}

func main() {
	user := User{Name: "sakura", Age: 0, Password: "123456"}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
