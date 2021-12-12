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
	value := ToUpper("mukesh")
	fmt.Println(value)
	if value != "MUKESH" {
		t.Fail()
	}

	value2 := ToUpper("MUkesh")
	fmt.Println(value2)
	if value2 != "MUKESH" {
		t.Errorf("Test Failed : want %v , got %v", "MUKESH", value2)
	}
}
