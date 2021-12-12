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
		if text[i] >= 97 && text[i] <= (97+26) {
			newString += (string(text[i] - (97 - 65)))
		} else {
			newString += (string(text[i]))
		}
	}
	return newString
}
