package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 1.死循环
	//for {
	//	fmt.Println("死循环")
	//}

	// 2.while循环
	sum1 := 0
	i := 1
	for i <= 100 {
		sum1 += i
		i++
	}
	fmt.Println(sum1)

	// 3.do while循环
	sum2 := 0
	i = 1
	for {
		sum2 += i
		i++
		if i > 100 {
			break
		}
	}
	fmt.Println(sum2)
}
