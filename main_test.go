package main

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	value := Greet("Git")
	fmt.Println(value)
}

func TestToUpper(t *testing.T) {
	value := ToUpper("vikas")
	fmt.Println(value)
	if value != "VIKAS" {
		t.Fail()
	}
}
