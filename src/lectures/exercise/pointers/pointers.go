//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func main() {
	//  - Create at least 4 items, all with active security tags
	tag1 := Item{
		name: "Hosptial",
		tag:  Active,
	}

	tag2 := Item{
		name: "Tag2",
		tag:  Active,
	}

	tag3 := Item{
		name: "Tag3",
		tag:  Active,
	}

	securityTags := []Item{tag1, tag2, tag3}
	fmt.Printf("Initial tags %+v \n", securityTags)
	securityTagPtr := &securityTags

	//  - Deactivate any one security tag in the array/slice
	setTag(&securityTags[1], Inactive)

	fmt.Printf("Item %+v has been desactivated\n", securityTags[1])
	fmt.Printf("Current tags %+v \n", securityTags)

	//  - Call the checkout() function to deactivate all tags
	checkout(securityTagPtr)

	//  Print out the array/slice after each change
	fmt.Printf("Last status tags %+v \n", securityTags)

}

func setTag(item *Item, tagStatus SecurityTag) {
	item.tag = tagStatus
}

func checkout(securityTags *[]Item) {
	items := *securityTags

	for key, _ := range items {
		items[key].tag = Inactive
	}
}
