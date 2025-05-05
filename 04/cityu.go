package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	city := "Krakow"
	fmt.Println(utf8.RuneCountInString(city))
}
