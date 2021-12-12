package main

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	value := Greet("Git")
	fmt.Println(value)
}
