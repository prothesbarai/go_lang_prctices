package main

import "fmt"

func multiplicationTable() {
	var number int
	output := 0
	fmt.Print("Give any Number : ")
	fmt.Scan(&number)

	for i := 1; i <= 10; i++ {
		output = number * i
		fmt.Printf("%dx%d=%d\n",number,i,output)
	}
}