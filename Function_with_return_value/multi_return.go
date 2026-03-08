package main

import (
	"fmt"
)

func multiReturn() {
	var n int
	fmt.Print("Enter Numbers Array Length : ")
	fmt.Scan(&n)
	numbers := make([]int, n) // Here make n size array/slice
	for i := 0; i < n; i++ {
		fmt.Printf("Enter Array Number Position-%d : ",i)
		fmt.Scan(&numbers[i])
	}

	min,max := getMinMax(numbers)
	fmt.Printf("In this array min number is %d & max number is %d",min,max)
}

func getMinMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}