package main
import "fmt"
func main(){
	var name string = "Prothes Barai"
	age := 27
	height := 5.6
	isDeveloper := false

	info := fmt.Sprintf("My name is %s and i'm %d. My height is : %.1f. And my all info is %t",name,age,height,isDeveloper);
	fmt.Println(info)
}