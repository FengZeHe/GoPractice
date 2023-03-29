package main

import "fmt"

type notifier interface {
	notify()
}

type myuser struct {
	name  string
	email string
}

func (u *myuser) notify() {
	fmt.Printf("Send user email to %s <%s>", u.name, u.email)
}

func main() {
	u := myuser{"feng", "feng@email.com"}
	sendNotification(&u)

}

func sendNotification(n notifier) {
	n.notify()
}
