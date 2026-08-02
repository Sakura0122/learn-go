package main

import "fmt"

func main() {
	// break用于跳出当前循环
	// continue用于跳过本轮循环

	// 99 乘法表
	//for i := 1; i <= 9; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%d*%d=%d ", j, i, i*j)
	//	}
	//	fmt.Println()
	//}

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j > i {
				break
			}
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
}
