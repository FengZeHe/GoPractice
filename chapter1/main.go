package main

import (
	_ "github.com/GoPractice/chapter1/matchers"
	search "github.com/GoPractice/chapter1/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	//main函数作为整个系统的入口
	search.Run()
}
