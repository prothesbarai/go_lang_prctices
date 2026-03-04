package main
import "fmt"
func basicForLoop(){
	count := 0
	fmt.Print("Enter Basic Loop Length : ")
	fmt.Scan(&count)
	for i := 1;  i<=count; i++ {
		fmt.Println("Iteration:", i)
	}
}



func loopWithConditon(){
	count := 0
	fmt.Print("Enter Condition Loop Even / odd Number Find : ")
	fmt.Scan(&count)
	for i := 1; i <= count; i++ {
		if (i%2 == 0) {
			fmt.Print(i," is Even\n")
		}else{
			fmt.Print(i," is Odd\n")
		}
	}
}

func miniProjectSimpleCalculator(){
	var fnum,lnum float64
	var operator string
	fmt.Print("Enter First Number : ")
	fmt.Scan(&fnum)
	fmt.Print("Enter Operator : ")
	fmt.Scan(&operator)
	fmt.Print("Enter Second Number : ")
	fmt.Scan(&lnum)

	switch operator {
	case "+":
		fmt.Print(fnum+lnum)
	case "-":
		fmt.Print(fnum-lnum)
	case "*":
		fmt.Print(fnum*lnum)
	case "/":
		if(lnum != 0){
			fmt.Print(fnum/lnum)
		}else{
			fmt.Print("You Cann't Devided By Zero!")
		}
	default:
		fmt.Print("Invalid Operator!")
	}

}