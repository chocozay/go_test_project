package main

import (
	"fmt"
	"math"
)

func square_root(num1 int) float64 {
	return math.Sqrt(num1)
}

func subtraction(num1 int, num2 int) int {
	return abs(num1 - num2)
}

func abs(num1 int) int {
	if num1 < 0 {
		return (num1 * -1)
	} else {
		return num1
	}
}

func multiplication(num1, num2 int) int {
	return num1 * num2
}

func logarithm(num1 float64) float64 {
	return math.Log10(num1)
}

func add(x int, y int) int {
	return x + y
}

func divide(x int, y int) int {
	return x / y
}

func percentage(x int, y int) int {
	return divide(x, y) * 100
}

func power(x int, y int) int {
	return math.Pow(x, y)
}

func main() {
	fmt.Println("the following ar the available functions")
	fmt.Println("Add, power, percentage, divide, substract, log, squareroot, and multiply")
	fmt.Print("Enter math operation or enter 'exit' to exit: ")
	var input string
	fmt.Scanln(&input)
	calculator(input)
}

func get_input() int {
	fmt.Println("enter value: ")
	var input int
	fmt.Scanln(&input)
	return input
}

func get_two_input() int {
	return [2]int{get_input(), get_input()}
}

func get_answer_with_one_input(method fn) {
	fmt.Println(method(get_input()))
}

func get_answer_with_two_input(method fn) {
	fmt.Println(method(get_two_input()))
}

func calculator(input string) float64 {
	switch input {
	case "add":
		get_answer_with_two_input(math_operations.add)
	case "power":
		get_answer_with_two_input(math_operations.power)
	case "divide":
		get_answer_with_two_input(math_operations.divide)
	case "percent":
		get_answer_with_two_input(math_operations.percentage)
	case "subtract":
		get_answer_with_two_input(math_operations.subtraction)
	case "squareroot":
		get_answer_with_one_input(math_operations.squareroot)
	case "log":
		get_answer_with_one_input(math_operations.logarithm)
	case "multiply":
		get_answer_with_two_input(math_operations.multiply)
	case "exit":
		fmt.Println("Thank you for using this calculator")
		exit(0)
	default:
		fmt.Println("Invalid input try again")
	}
}
