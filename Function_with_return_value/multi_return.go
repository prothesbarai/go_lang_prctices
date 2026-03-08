package main

import (
	"fmt"
)

func multiReturn() {
	var num int
	fmt.Print("Enter Size of Number List : ")
	fmt.Scan(&num)
	numbers :=  make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Printf("Enter Number Position-%d : ",i)
		fmt.Scan(&numbers[i])
	}


	min,max := multiReturnFuncTask(numbers)
	fmt.Printf("Min Number is : %d, And max number is %d",min,max)

}

func multiReturnFuncTask(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers{
		if(num < min){
			min = num
		}
		if(num > max){
			max = num
		}
	}

	return min, max
}