//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Products struct {
	Price float64
	Name  string
}

var ShoppingList [4]Products

func main() {
	ShoppingList[0] = Products{Price: 50, Name: "Javon"}
	ShoppingList[1] = Products{Price: 35.50, Name: "Manzana"}
	ShoppingList[2] = Products{Price: 20, Name: "Pescado"}

	lastIndex := len(ShoppingList) - 1
	fmt.Println(lastIndex)
	fmt.Println("The last item on the list: ", ShoppingList[lastIndex])
	fmt.Println("The total number of items: ", len(ShoppingList))
	var sum float64

	for i := 1; i < len(ShoppingList); i++ {
		sum += ShoppingList[i].Price
	}
	fmt.Printf("The total number of items: %v \n", sum)

}
