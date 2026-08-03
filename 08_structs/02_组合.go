package main

import "fmt"

type Class struct {
	Name string
}

type Student struct {
	Class
	Name string
	Age  int
}

func (s *Student) getInfo() {
	fmt.Println(s.Name, s.Age, s.Class.Name)
}

func (s *Student) setName(name string) {
	s.Name = name
}

func main() {
	s1 := Student{Class: Class{Name: "A"}, Name: "sakura", Age: 24}
	s1.getInfo()

	s2 := Student{}
	s2.setName("sakura")
	s2.getInfo()
}
