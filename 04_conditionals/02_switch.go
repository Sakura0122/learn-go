package main

import "fmt"

func main() {
	var age int

	fmt.Print("请输入你的年龄：")
	fmt.Scan(&age)

	switch {
	case age > 18:
		fmt.Println("你已经成年了")
	case age > 0:
		fmt.Println("你还未成年")
	default:
		fmt.Println("输入的年龄无效")
	}

	var week int

	fmt.Print("请输入星期几：")
	fmt.Scan(&week)
	switch week {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	default:
		fmt.Println("输入的星期无效")
	}
}
