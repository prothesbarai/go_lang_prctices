package main
import "fmt"
func nestedElseIf(){
	marks := 0.0
	fmt.Print("Enter Your Marks : ")
	fmt.Scan(&marks)
	if(marks >= 80 && 100 >= marks){
		fmt.Print("You Got A+")
	}else if(marks >= 70 && 80 > marks){
		fmt.Print("You Got A")
	}else if(marks >= 60 && 70 > marks){
		fmt.Print("You Got A-")
	}else if(marks >= 50 && 60 > marks){
		fmt.Print("You Got B")
	}else if(marks >= 40 && 50 > marks){
		fmt.Print("You Got C")
	}else if(marks >= 35 && 40 > marks){
		fmt.Print("You Got D")
	}else if(marks >= 33 && 35 > marks){
		fmt.Print("You are Passed")
	}else if(marks < 33 && 0 <= marks){
		fmt.Print("You are Failed")
	}else{
		fmt.Print("Invalid Marks")
	}
}