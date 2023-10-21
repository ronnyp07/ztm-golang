//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	for i := 0; i <= 50; i++ {

		if devidedByThree(i) && devidedByFive(i) {
			fmt.Println("FizzBuzz: ", i)
		} else if devidedByThree(i) {
			fmt.Println("Fizz: ", i)
		} else if devidedByFive(i) {
			fmt.Println("Buzz: ", i)
		} else {
			fmt.Println(i)
		}

	}
}

func devidedByThree(num int) bool {
	return num%3 == 0
}

func devidedByFive(num int) bool {
	return num%5 == 0
}
