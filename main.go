package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Greet("Git"))
}

func Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

func ToUpper(text string) string {
	return strings.ToUpper(text)
}
