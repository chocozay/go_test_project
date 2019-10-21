package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func divide(x int, y int) int {
	return x / y
}

func percentage(x int, y int) int {
	return divide(x, y) * 100
}

func main() {
	fmt.Println("the following ar the available functions")
	fmt.Println("Add, power, percentage, division")
	fmt.Print("Enter math operation: ")
	var input string
	fmt.Scanln(&input)

	switch input {
	case "add":
		fmt.Println("enter value 1")
		var input1 int
		fmt.Scanln(&input1)
		fmt.Println("enter value 2")
		var input2 int
		fmt.Scanln(&input2)
		fmt.Println(add(input1, input2))
	case "power":
		fmt.Println("enter value 1")
		var input1 float64
		fmt.Scanln(&input1)
		fmt.Println("enter value 2")
		var input2 float64
		fmt.Scanln(&input2)
		fmt.Println(math.Pow(input1, input2))
	case "divide":
		fmt.Println("enter value 1")
		var input1 int
		fmt.Scanln(&input1)
		fmt.Println("enter value 2")
		var input2 int
		fmt.Scanln(&input2)
		fmt.Println(divide(input1, input2))
	case "percent":
		fmt.Println("enter value 1")
		var input1 int
		fmt.Scanln(&input1)
		fmt.Println("enter value 2")
		var input2 int
		fmt.Scanln(&input2)
		fmt.Println(percentage(input1, input2))
	}
}
