package main

import "fmt"

func add(num1, num2 int) int {
	return num1 + num2
}

func compute(num1, num2 int, op func(num1, num2 int) int) int {
	return op(num1, num2)
}

func main() {
	result := compute(2, 2, add)
	fmt.Printf("2 + 2: %v\n", result)
}
