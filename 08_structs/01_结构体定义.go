package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) study() {
	fmt.Println(s.Name, "正在学习")
}

func main() {
	s1 := Student{Name: "sakura", Age: 24}
	s1.study()

	s2 := Student{"zs", 25}
	s2.study()
}
