package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	fmt.Println("Create goroutines")
	go PrintPrime("A")
	go PrintPrime("B")

	//	等待goroutine结束
	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("Terminating Program")

}

func PrintPrime(prefix string) {
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)

}
