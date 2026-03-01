package main
import "fmt"
func simpleUserInputProgram(){
	name := ""
	age := 0
	salary := 0.0
	fmt.Print("Name : ")
	fmt.Scan(&name)
	fmt.Print("Age : ")
	fmt.Scan(&age)
	fmt.Print("Salary : ")
	fmt.Scan(&salary)

	output := fmt.Sprintf("Hello %s\nYour age is %d\nYour yearly salary is %0.f",name,age,salary)
	fmt.Print(output)
}