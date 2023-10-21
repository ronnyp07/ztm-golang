//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	Length int
	Width  int
}

func main() {
	rect1 := Rectangle{4, 10}
	fmt.Println(Area(rect1.Length, rect1.Width))

	var rect2 Rectangle
	rect2.Length = 5
	rect2.Width = 10

	fmt.Println(Area(rect2.Length, rect2.Width))
	fmt.Println("The Area is: ", Area2(rect2))

	fmt.Println("Area of the first rectangle: ", rect1.CalculateArea())

}

func Area(length, width int) int {
	return length * width
}

func Area2(item Rectangle) int {
	return item.Length * item.Width
}

func (r Rectangle) CalculateArea() int {
	return r.Length * r.Width
}
