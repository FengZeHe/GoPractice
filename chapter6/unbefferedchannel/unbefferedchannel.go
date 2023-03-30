package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 创建一个无缓冲的通道
	court := make(chan int)

	wg.Add(2)
	go player("Nadal", court)
	go player("Dj", court)

	court <- 1
	wg.Wait()
}

// 模拟一个选手在打网球
func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won \n", name)
			return
		}

		n := rand.Intn(10)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed \n", name)

			//关闭通道 表示输了
			close(court)
			return
		}

		// 显示击球数
		fmt.Printf("Player %s Hit %d \n", name, ball)
		ball++

		court <- ball
	}
}
