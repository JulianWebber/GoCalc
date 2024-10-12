package main

import (
	"fmt"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b != 0 {
		return a / b
	}
	fmt.Println("Error: Division by zero.")
	return 0
}

func main() {
	var a, b float64
	var op string

	fmt.Println("Enter first number:")
	fmt.Scanln(&a)
	fmt.Println("Enter second number:")
	fmt.Scanln(&b)
	fmt.Println("Enter operation (+, -, *, /):")
	fmt.Scanln(&op)

	switch op {
	case "+":
		fmt.Printf("Result: %.2f\n", add(a, b))
	case "-":
		fmt.Printf("Result: %.2f\n", subtract(a, b))
	case "*":
		fmt.Printf("Result: %.2f\n", multiply(a, b))
	case "/":
		result := divide(a, b)
		if b != 0 {
			fmt.Printf("Result: %.2f\n", result)
		}
	default:
		fmt.Println("Invalid operation. Please enter one of +, -, *, /.")
	}
}