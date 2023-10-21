package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Size int

const (
	Unkowned Size = iota
	Small
	Large
)

func (s *Size) UnmarshalText(text []byte) error {

	switch strings.ToLower(string(text)) {
	case "small":
		*s = Small
	case "large":
		*s = Large
	default:
		*s = Unkowned
	}

	return nil
}

func (s Size) MarshalText() ([]byte, error) {
	var name string
	switch s {
	case Small:
		name = "small"
	case Large:
		name = "large"
	default:
		name = "unkowned"
	}
	return json.Marshal(name)
}

func main() {
	blob := `["small","regular","large","unrecognized","small","normal","small","large"]`
	result := make([]Size, 0)
	if err := json.Unmarshal([]byte(blob), &result); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	fmt.Println()

	unmarshalResult, _ := json.Marshal(result)
	fmt.Println(string(unmarshalResult))

}
