package main
import "fmt"

func circleAreaCalculation(){
	radius := 0.0
	fmt.Print("Circle Radius : ")
	fmt.Scan(&radius)
	area := 3.1416 * radius * radius
	output := fmt.Sprintf("Area = %.2f",area)
	fmt.Print(output)
}