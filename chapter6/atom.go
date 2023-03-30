package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	atom_counter int64
	atom_wg      sync.WaitGroup
)

func main() {
	atom_wg.Add(2)
	go atom_incCounter(1)
	go atom_incCounter(2)
	atom_wg.Wait()
	fmt.Println("finfish", atom_counter)
}

func atom_incCounter(id int) {
	defer atom_wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&atom_counter, 1)

		runtime.Gosched()
	}

}
