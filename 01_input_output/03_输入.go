package main

import "fmt"

func main() {
	fmt.Print("请输入你的名字：")

	var name string

	fmt.Scan(&name)
	fmt.Printf("%s\n", name)

	fmt.Print("请输入你的年龄：")
	var age int
	// n 是 fmt.Scan 成功读取并赋值的参数个数。
	n, err := fmt.Scan(&age)
	fmt.Println(n, err, age)
}
