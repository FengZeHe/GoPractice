package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 创建一个无缓冲头通道
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	baton <- 1
	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int
	runner := <-baton

	// 开始跑步
	fmt.Printf("Runner %d Running With Baton \n", runner)

	//创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To the Line \n", newRunner)
		go Runner(baton)
	}

	// 模拟围绕跑道跑
	time.Sleep(100 * time.Millisecond)

	// 比赛结束了吗？
	if runner == 4 {
		fmt.Printf("Runner %d Finfished ,Race Over \n", runner)
		wg.Done()
		return
	}

	// 将接力棒交给下一位跑步者
	fmt.Printf("Runner %d exchange with runner %d \n", runner, newRunner)

	baton <- newRunner
}
