package main

import "fmt"

func main() {
	fmt.Println(Greet("Git"))
}

func Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
