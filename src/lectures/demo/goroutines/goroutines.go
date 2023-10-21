package main

import (
	"fmt"
	"time"
	"unicode"
)

func Capitalletter(item rune, upperRune *[]rune) {
	*upperRune = append(*upperRune, unicode.ToUpper(item))
}

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	upperRun := make([]rune, 0)

	for i := 0; i < len(data); i++ {
		go Capitalletter(data[i], &upperRun)
	}

	time.Sleep(5 * time.Second)

	fmt.Printf("%c\n", upperRun)

}
