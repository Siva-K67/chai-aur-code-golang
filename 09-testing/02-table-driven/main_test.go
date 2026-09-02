package main

import "testing"

//What if you want to test several input/output pairs without
//writing a separate if block for each? This is where table-driven tests
//come in — a very common Go pattern.

func TestAdd(t *testing.T) {
	// a "table" - a slice of structs, each holding one test case
	testCases := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{10, -5, 5},
	}

	// loop through the table, running the same check for every case
	for _, tc := range testCases {
		result := add(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("add(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}
