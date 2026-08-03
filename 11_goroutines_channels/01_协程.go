package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wait = sync.WaitGroup{}
)

func sing() {
	fmt.Println("唱歌")
	time.Sleep(1 * time.Second)
	fmt.Println("唱歌结束")
	wait.Done()
}

func main() {
	startTime := time.Now()

	wait.Add(3)
	go sing()
	go sing()
	go sing()
	wait.Wait()

	fmt.Println("main结束", time.Since(startTime))
}
