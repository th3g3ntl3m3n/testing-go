package main

import (
	"fmt"
)

func main() {
	fmt.Println(Greet("Git"))
}

func Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

func ToUpper(text string) string {
	var newString = ""
	for i := 0; i < len(text); i++ {
		newString += (string(text[i] - (97 - 65)))
	}
	return newString
}
