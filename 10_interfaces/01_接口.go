package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + ": 汪"
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return p.Name + ": 你好"
}

func PrintSpeech(s Speaker) {
	_, ok := s.(Dog)
	if ok {
		fmt.Println("是 Dog 类型")
	} else {
		fmt.Println("不是 Dog 类型")
	}
	fmt.Println(s.Speak())
}

func main() {
	PrintSpeech(Dog{Name: "小黑"})
	PrintSpeech(Person{Name: "小明"})
}
