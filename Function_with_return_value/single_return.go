package main

import "fmt"

func singleReturnFunc() {
	var num1 float64
	var num2 float64
	fmt.Print("Enter 1st Number : ")
	fmt.Scan(&num1)
	fmt.Print("Enter 2nd Number : ")
	fmt.Scan(&num2)

	result := singleReturnFuncTask(num1, num2)

	fmt.Print("Final Output : ",result)
}

func singleReturnFuncTask(num1 float64, num2 float64) float64 {
	return num1 + num2
}