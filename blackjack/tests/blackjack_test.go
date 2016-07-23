package tests

import "testing"

func TestAreaofSquare(t *testing.T) {
	s := Square{side: 2}
	result := s.Area()

	if result != 4 {
		t.Errorf("Expected area of square to be 4, got %d", result)
	}
}

