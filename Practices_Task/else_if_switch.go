package main
import "fmt"

func positiveNegativeNumberCheck(){
	number := 0.0
	fmt.Print("Give a Number : ")
	fmt.Scan(&number)
	if(number != 0){
		if(number > 0){
			fmt.Print(number," It's Positive Number")
		}else{
			fmt.Print(number," It's Negative Number")
		}
	}else{
		fmt.Print(number," It's Zero")
	}
}



func monthChecker(){
	month := 0
	fmt.Print("Get Month Position Number : ")
	fmt.Scan(&month)
	switch month {
	case 1:
		fmt.Print("January")
	case 2:
		fmt.Print("February")
	case 3:
		fmt.Print("March")
	case 4:
		fmt.Print("April")	
	case 5:
		fmt.Print("May")
	case 6:
		fmt.Print("June")
	case 7:
		fmt.Print("July")
	case 8:
		fmt.Print("Augest")
	case 9:
		fmt.Print("September")
	case 10:
		fmt.Print("October")
	case 11:
		fmt.Print("November")
	case 12:
		fmt.Print("December")		
	default:
		fmt.Print("Invalid Month Value")
	}
}