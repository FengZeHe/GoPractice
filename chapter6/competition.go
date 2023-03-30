package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	mywg    sync.WaitGroup
)

func main() {
	mywg.Add(2)
	go incCounter(1)
	go incCounter(2)

	mywg.Wait()
	fmt.Println("finfish -> counter=", counter)
}

func incCounter(id int) {
	defer mywg.Done()

	for count := 0; count < 2; count++ {
		value := counter

		// 当前goroutine从线程退出，并放回队列
		runtime.Gosched()

		value++
		counter = value
	}

}
