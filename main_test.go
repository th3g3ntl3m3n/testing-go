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

// feature - [ToUpper, ToCamel]

// staging - [staging.swiftbit.xyz] - [experimental] -/ [nightly, experimental]

// testing - [testing.swiftbit.xyz] - [ready to test] -

// [release - 0.0.1, 0.0.2, 0.0.3, 0.0.4, 0.1.0, 0.1.1]

// production - 0.1.0 [admin.swiftbit.xyz]
