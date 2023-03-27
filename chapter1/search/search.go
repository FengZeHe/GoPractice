package search

import (
	"fmt"
	"math/rand"
	"sync"
)

//var matchers = make(map[string]Matcher)

func init() {
	fmt.Println("search init")
}

func Run() {

	var waitGroup sync.WaitGroup
	myarr := []int{1, 2, 3, 4, 5, 6}
	waitGroup.Add(len(myarr))

	for _, v := range myarr {
		go func(v int) {
			res, _ := calculation(v)
			fmt.Println(res)
			waitGroup.Done()
		}(v)
	}

	waitGroup.Wait()

	fmt.Println("this is search.Run")
}

func calculation(number int) (int, error) {
	res := number + rand.Int()
	return res, nil
}
