package main

import (
	"fmt"
)
type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	user1 := User{Name: "Prothes", Age: 12, Email: "email"}
	fmt.Print(user1)
	fmt.Printf("%T",user1)
}
