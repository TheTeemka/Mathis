package logic

import (
	"testing"
)

func TestSolveExpr(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"10 * 24 - 24", 216},
		{"-(10 * 5) + 23", -27},
		{"(43 + 27) * (5 - 3)", 140},
		{"-(23) * 12", -276},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := SolveExpr(tt.input)
			if result != tt.expected {
				t.Errorf("SolveExpr(%s) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
