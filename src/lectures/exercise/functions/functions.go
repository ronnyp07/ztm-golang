//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num1, num2, num3 int64 = 5, 7, 10
	fmt.Println(greeting("Ronny"))
	fmt.Println(anyNumber())
	fmt.Println(twoRandomNum())
	fmt.Println(fmt.Sprintf("The result of %v+%v+%v=%v", num1, num2, num3, sum(num1, num2, num3)))
}

func greeting(name string) string {
	return "Hello " + name
}

func sum(num1 int64, num2 int64, num3 int64) int64 {
	return num1 + num2 + num3
}

func anyNumber() float64 {
	return rand.Float64()
}

func twoRandomNum() (float64, float64) {
	return anyNumber(), anyNumber()
}
