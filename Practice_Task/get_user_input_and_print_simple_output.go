package main
import "fmt"
func main(){
	name := ""
	age := 0
	salary := 0.0
	fmt.Print("Name : ")
	fmt.Scan(&name)
	fmt.Print("Age : ")
	fmt.Scan(&age)
	fmt.Print("Salary : ")
	fmt.Scan(&salary)
	info := fmt.Sprintf("Hello %s\nYour age is %d\nYour yearly salary is %.1f",name,age,salary)
	fmt.Print(info)
}