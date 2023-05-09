package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5.5, 5)
	if total != 10.5 {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", total, 10.5)
	}
}
