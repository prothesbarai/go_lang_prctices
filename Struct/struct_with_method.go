package main

import "fmt"

type User2 struct {
	Name string
	Age  int
}

func (u User2) Greet() string {
	return "Hello " + u.Name
}

func structWithMethod() {
	user1 := User2{Name: "Prothes", Age: 12}
	fmt.Print(user1.Greet())
}