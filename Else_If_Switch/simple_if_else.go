package main
import "fmt"

func simpleIfElse(){
	age := 0.0
	fmt.Print("Enter Your Age : ")
	fmt.Scan(&age)
	if (age >= 18.0) {
		fmt.Print("You are Adult Person")
	}else{
		fmt.Print("You are Child")
	}
}