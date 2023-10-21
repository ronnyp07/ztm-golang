package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Animal int

const (
	Unkown Animal = iota
	Gopher
	Zebra
)

func (a *Animal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	default:
		*a = Unkown
	}
	fmt.Println("unmarshal called for: ", s)
	return nil
}

func (a Animal) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	default:
		s = "unkown"

	}

	return json.Marshal(s)
}

func main() {
	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
	var zoo []Animal
	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
		log.Fatal(err)
	}

	groupZoo := make(map[Animal]int)

	for _, item := range zoo {
		groupZoo[item] += 1
	}

	fmt.Println("The total of Gopher is: ", groupZoo[Gopher])
	fmt.Println("The total of Zebra is: ", groupZoo[Zebra])
	fmt.Println("The total of Unknow is: ", groupZoo[Unkown])

	blob2 := []Animal{Unkown, Zebra, Zebra, Gopher}
	result, _ := json.Marshal(blob2)

	fmt.Println("A list of animals: ", string(result))
}
