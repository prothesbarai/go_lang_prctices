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
	
}