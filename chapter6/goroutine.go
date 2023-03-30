package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// wg 用来等待程序完成
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()
		for count := 0; count < 2; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf(" %c ", char)
			}
		}
		fmt.Println(" ")
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 2; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf(" %c ", char)
			}
		}
		fmt.Println(" ")

	}()

	fmt.Println("Waiting to finish \n")
	wg.Wait()
	fmt.Println("\n Done ...")

}
