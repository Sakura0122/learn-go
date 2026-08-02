package main

import "fmt"

func main() {
	var age int

	fmt.Print("请输入你的年龄：")
	fmt.Scan(&age)
	if age > 18 {
		fmt.Println("你已经成年了")
	} else if age > 0 {
		fmt.Println("你还未成年")
	} else {
		fmt.Println("输入的年龄无效")
	}
}
