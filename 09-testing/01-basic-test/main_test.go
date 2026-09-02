package main

import "testing"

// test functions must start with "Test", and take *testing.T as their only parameter
// Go's test runner literally scans for functions matching this naming pattern.
func TestAdd(t *testing.T) {
	result := add(2, 3)
	expected := 5

	if result != expected {
		// t.Errorf marks this test as FAILED, but keeps running other tests
		t.Errorf("add(2, 3) = %d; expected %d", result, expected)
	}
}
