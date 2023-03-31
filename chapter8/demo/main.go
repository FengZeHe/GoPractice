package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	//创建一个buffer值，并将一个字符串写入buffer
	var b bytes.Buffer
	b.Write([]byte("hello"))

	// 字符串拼接
	fmt.Fprintf(&b, "World!")

	b.WriteTo(os.Stdout)
}
