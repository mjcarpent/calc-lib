package calc

import (
	"fmt"
	"testing"
)

func TestAddition_Calculate(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name       string
		a, b, want int
	}{
		{"Zero as the second operand: ", 1, 0, 1},
		{"Zero as the first operand: ", 0, 1, 1},
		{"Basic int for first and second operands: ", 4, 1, 5},
		{"Negative int for first operand end with positive: ", -5, 6, 1},
		{"Negative int for second operand end with positive: ", 7, -6, 1},
		{"Negative int for first operand end with negative: ", -5, 4, -1},
		{"Negative int for second operand end with negative: ", 7, -8, -1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v: %d + %d = %d", tt.name, tt.a, tt.b, tt.want)
		t.Run(testname, func(t *testing.T) {
			this := &Addition{}
			if got := this.Calculate(tt.a, tt.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
