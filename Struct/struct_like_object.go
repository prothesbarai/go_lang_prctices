package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func structDataAttributes() {
	user1 := User{
		ID: 1, Name: "Prothes", Email: "developerprothes16@gmail.com", Age: 28,
	}
	fmt.Println("User Name : ",user1.Name)
	fmt.Println("User Email : ",user1.Email)
}
