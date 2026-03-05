package main

import "fmt"

func factorialNumberCalculate() {
	var number int
	fact := 1
	fmt.Print("Enter Given Number : ")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		fact *= i
	}

	fmt.Printf("Factorial of %d is %d\n",number,fact)
}