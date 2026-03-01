package main
import "fmt"
func switchCase(){
	day := 0
	fmt.Print("Enter Day Value : ")
	fmt.Scan(&day)

	switch (day) {
	case 1:
		fmt.Print("Monday")
	case 2:
		fmt.Print("Twesday")
	case 3:
		fmt.Print("Wednesday")
	case 4:
		fmt.Print("Thusday")
	case 5:
		fmt.Print("Friday")
	case 6:
		fmt.Print("Saturday")
	case 7:
		fmt.Print("Sunday")
	default:
		fmt.Print("Invalid Day Value")	
	}

}