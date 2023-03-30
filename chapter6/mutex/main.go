package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println("counter", counter)
}

func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个goroutine进入 这个临界区
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()

	}
}
