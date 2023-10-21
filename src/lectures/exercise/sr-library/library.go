//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import "time"

type Title string //book title
type Name string  // library member name

type Book struct {
	title    string
	isBooked bool
}

type LendAudit struct {
	checkout time.Time
	checkIn  time.Time
}

type Members struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int
	lended int
}

type Library struct {
	member map[Name]Member
	books  map[Title]BookEntry
}

func main() {
	book1 := 
}
