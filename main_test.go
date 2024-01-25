package main

import "testing"

func TestDivide(t *testing.T) {
	_, err := divide(100, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}
