package main

import "testing"

var tests = []struct {
	name     string
	divided  float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0, 0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
	{"expect-zero", 0.0, 10.0, 0.0, false},
	{"expect-negative", -100.0, 10.0, -10.0, false},
	{"expect-large-number", 1000000.0, 1000.0, 1000.0, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.divided, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(100, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}
