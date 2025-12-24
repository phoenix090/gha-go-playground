package mathutil

import "testing"

// return the sum of two integers
func TestAdd(t *testing.T) {
	result := add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("add(2, 3) = %d; want %d", expected, result)
	}
}

func TestAddNegative(t *testing.T) {
	result := add(-2, -3)
	expected := -5
	if result != expected {
		t.Errorf("add(-2, -3) = %d; want %d", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("subtract(5, 3) = %d; want %d", expected, result)
	}
}
