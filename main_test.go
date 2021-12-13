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
		{value: "MUkesH", expect: "MUKESH"},
	}

	for _, td := range testData {
		got := ToUpper(td.value)
		if got != td.expect {
			t.Errorf("Error TC: %s, Expected %s, Got %s ", td.value, td.expect, got)
		}
	}
}
