package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int)

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 支付 %d 元\n", name, money)
	time.Sleep(1 * time.Second)
	moneyChan <- money
	wait.Done()
}

func main() {
	var wait sync.WaitGroup
	wait.Add(2)
	go pay("张三", 100, &wait)
	go pay("李四", 200, &wait)
	go func() {
		wait.Wait()
		close(moneyChan)
	}()

	var moneyList []int
	for money := range moneyChan {
		fmt.Println("收到支付金额：", money)
		moneyList = append(moneyList, money)
	}
	fmt.Println("支付完成，总金额为：", moneyList)
}
