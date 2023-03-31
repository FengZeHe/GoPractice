package main

import "log"

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// 写到标准日志记录器
	log.Println("message")

	// Fatalln 在调用Println之后接着调用 os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln 调用之后会调用panic
	log.Panicln("panic message")

}
