package main

import "fmt"

type user struct {
	name string
	age  int
	sex  string
}

type school struct {
	person user
	id     string
}

func RunType() {
	var student user
	student.name = "feng"
	student.age = 19
	student.sex = "male"

	teacher := user{
		name: "teacher",
		age:  66,
		sex:  "female",
	}

	tim := school{id: "123456", person: user{age: 18, name: "Tim", sex: "female"}}

	fmt.Println(student, teacher, tim)
}

func main() {
	RunType()
}
