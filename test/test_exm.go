package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter Size of Integer Array : ")
	fmt.Scan(&n)
	var numbers = make([]int,n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter Number Position-%d : ",i)
		fmt.Scan(&numbers[i])
	}
	max,min := finding(numbers)

	fmt.Printf("Maximum : %d\nMinimum : %d",max,min)
}


func finding(maxMin []int)(int,int){
	max := maxMin[0]
	min := maxMin[0]

	for _, num:=range maxMin {
		if(num > max){
			max = num
		}
		if(num < min){
			min = num
		}
	}

	return max,min
}