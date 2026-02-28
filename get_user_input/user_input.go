package main
import "fmt"
func main(){
	name := ""
	fmt.Print("Enter Your Name : ")
	fmt.Scan(&name)
	info := fmt.Sprintf("Name is %s",name)
	fmt.Print(info)
}
