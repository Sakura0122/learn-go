package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup
var mp = sync.Map{}
var stop = make(chan struct{}) // 添加停止信号

func reader() {
	defer wait.Done()
	for i := 0; i < 10; i++ { // 只读10次
		if val, ok := mp.Load("time"); ok {
			fmt.Println(val)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func writer() {
	defer wait.Done()
	for i := 0; i < 10; i++ { // 只写10次
		mp.Store("time", time.Now().Format("15:04:05"))
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	wait.Add(2)
	go reader()
	go writer()

	time.Sleep(5 * time.Second) // 运行5秒
	close(stop)                 // 发送停止信号
	wait.Wait()
}
