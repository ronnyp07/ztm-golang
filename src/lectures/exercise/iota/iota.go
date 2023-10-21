//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

const (
	Add Operation = iota
	Subtract
	Multiply
	Divide
)

type Operation int

func (d Operation) calculate(num, den int) int {
	switch d {
	case Add:
		return num + den
	case Subtract:
		return num - den
	case Multiply:
		return num * den
	case Divide:
		return num / den
	}

	return 0
}

func main() {
	add := Add
	fmt.Println(add.calculate(2, 2)) // = 4

	sub := Subtract
	fmt.Println(sub.calculate(10, 3)) // = 7

	mul := Multiply
	fmt.Println(mul.calculate(3, 3)) // = 9

	div := Divide
	fmt.Println(div.calculate(100, 2)) // = 50
}
