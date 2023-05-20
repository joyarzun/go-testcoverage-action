package main

import "testing"

func TestSum(t *testing.T) {
	calc := new(Calculator)
	want := 3
	got := calc.add(1, 2)

	if got != want {
		t.Errorf("Result was incorrect, got: %d, want: %d", got, want)
	}
}
