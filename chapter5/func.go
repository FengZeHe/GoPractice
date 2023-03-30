package main

import "fmt"

type users struct {
	name  string
	email string
}

func (u users) notify() {
	fmt.Printf("Sending User Email to %s<%s> \n", u.name, u.email)
}

// 使用指针修改一个方法
func (u *users) changeEmail(email string) {
	u.email = email
}

func (u users) hello() {
	fmt.Println("hello", u.name)
}

func main() {
	tim := users{"Tim", "Tim@email.com"}
	tim.notify()
	tim.hello()

	cock := &users{"cock", "cock@email.com"}
	cock.notify()
	cock.changeEmail("change@email.com")
	cock.notify()
}
