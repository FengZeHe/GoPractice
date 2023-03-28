package main

import "fmt"

func main() {
	arr := [...]int{1: 2, 4: 5}

	fmt.Println(arr)
	var arr1 [3]*string
	arr2 := [3]*string{new(string), new(string), new(string)}

	*arr2[0] = "Red"
	*arr2[1] = "Blue"
	*arr2[2] = "green"

	arr1 = arr2
	fmt.Println(arr1)

}
