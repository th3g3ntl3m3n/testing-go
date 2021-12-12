package main

import (
	"testing"
)

type TD struct {
	value  string
	expect string
}

func TestGreet(t *testing.T) {
	value := Greet("Git")
	if value != "Hello Git!" {
		t.Error("Error in here")
	}
}

func Test(t *testing.T) {
	testData := []TD{
		{value: "mukesh", expect: "MUKESH"},
		{value: "vikas", expect: "VIKAS"},
		{value: "Vikas Sharma", expect: "VIKAS SHARMA"},
		{value: "MUkesh", expect: "MUKESH"},
	}

	for _, td := range testData {
		got := ToUpper(td.value)
		if got != td.expect {
			t.Errorf("Error %s | %s != %s ", td.value, td.expect, got)
		}
	}
}
