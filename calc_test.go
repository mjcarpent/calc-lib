package calc

import (
	"math"
	"testing"
)

func TestAddition_Calculate(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Positive add a positive equals positive", args: args{1, 2}, want: 3},
		{name: "Positive add a negative equals zero", args: args{1, -1}, want: 0},
		{name: "Negative add a positive equals zero", args: args{-1, 1}, want: 0},
		{name: "Positive add a negative equals negative", args: args{1, -2}, want: -1},
		{name: "Negative add a positive equals negative", args: args{-2, 1}, want: -1},
		{name: "Positive add a negative equals positive", args: args{11, -2}, want: 9},
		{name: "Negative add a positive equals positive", args: args{-2, 11}, want: 9},
		{name: "Max int add a positive wraps negative", args: args{math.MaxInt, 1}, want: math.MinInt},
		{name: "Max int add a negative equals positive", args: args{math.MaxInt, -1}, want: math.MaxInt - 1},
		{name: "Min int add a positive equals negative", args: args{math.MinInt, 1}, want: math.MinInt + 1},
		{name: "Min int add a negative wraps positive", args: args{math.MinInt, -1}, want: math.MaxInt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Addition{}
			if got := this.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtraction_Calculate(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Positive minus a positive equals a positive", args: args{2, 1}, want: 1},
		{name: "Positive minus a positive equals zero", args: args{2, 2}, want: 0},
		{name: "Positive minus a positive equals a negative", args: args{2, 4}, want: -2},
		{name: "Positive minus a negative equals positive", args: args{1, -1}, want: 2},
		{name: "Negative minus a positive equals a negative", args: args{-2, 1}, want: -3},
		{name: "Negative minus a negative equals a negative", args: args{-2, -1}, want: -1},
		{name: "Negative minus a negative equals zero", args: args{-2, -2}, want: 0},
		{name: "Negative minus a negative equals a positive", args: args{-2, -3}, want: 1},
		{name: "Max int minus a positive equals a positive", args: args{math.MaxInt, 1}, want: math.MaxInt - 1},
		{name: "Max int minus a negative wraps negative", args: args{math.MaxInt, -1}, want: math.MinInt},
		{name: "Min int minus a positive wraps positive", args: args{math.MinInt, 1}, want: math.MaxInt},
		{name: "Min int minus a negative equals a negative", args: args{math.MinInt, -1}, want: math.MinInt + 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Subtraction{}
			if got := this.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplication_Calculate(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Positive times a positive equals a positive", args: args{2, 2}, want: 4},
		{name: "Positive times 1 equals the positive", args: args{2, 1}, want: 2},
		{name: "Positive times 0 equals zero", args: args{2, 0}, want: 0},
		{name: "Positive times a negative equals a negative", args: args{2, -2}, want: -4},
		{name: "Negative times a negative equals a positive", args: args{-2, -2}, want: 4},
		{name: "Negative times 1 equals the negative", args: args{-2, 1}, want: -2},
		{name: "Negative times 0 equals zero", args: args{-2, 0}, want: 0},
		{name: "Negative times a positive equals a negative", args: args{-2, 2}, want: -4},
		{name: "Max int times a positive wraps", args: args{math.MaxInt, 2}, want: -2},
		{name: "Max int times a negative equals negative", args: args{math.MaxInt, -2}, want: 2},
		{name: "Min int times a positive ", args: args{math.MinInt, 2}, want: 0},
		{name: "Min int times a negative wraps ", args: args{math.MinInt, -2}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Multiplication{}
			if got := this.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision_Calculate(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Positive divide a positive equals a positive", args: args{2, 2}, want: 1},
		{name: "Positive divide 1 equals the positive", args: args{2, 1}, want: 2},
		{name: "Positive divide a negative equals a negative", args: args{2, -2}, want: -1},
		{name: "Negative divide a negative equals a positive", args: args{-2, -2}, want: 1},
		{name: "Negative divide 1 equals the negative", args: args{-2, 1}, want: -2},
		{name: "Negative divide 0 equals zero", args: args{0, -2}, want: 0},
		{name: "Negative divide a positive equals a negative", args: args{-2, 2}, want: -1},
		{name: "Max int divide a positive equals a positive", args: args{math.MaxInt, int(1000000000000000000)}, want: 9},
		{name: "Max int divide a negative equals a negative", args: args{math.MaxInt, int(-1000000000000000000)}, want: -9},
		{name: "Min int divide a positive equals a negative", args: args{math.MinInt, int(1000000000000000000)}, want: -9},
		{name: "Min int divide a negative equals a positive", args: args{math.MinInt, int(-1000000000000000000)}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Division{}
			if got := this.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModulus_Calculate(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Positive mod a positive equals a positive", args: args{2, 2}, want: 0},
		{name: "Positive mod 1 equals the positive", args: args{3, 2}, want: 1},
		{name: "Positive mod a negative equals a negative", args: args{3, -2}, want: 1},
		{name: "Negative mod a negative equals a positive", args: args{-3, -2}, want: -1},
		{name: "Negative mod 1 equals zero", args: args{-2, 1}, want: 0},
		{name: "Negative mod 0 equals negative", args: args{0, -2}, want: 0},
		{name: "Negative mod a positive equals a negative", args: args{-2, 3}, want: -2},
		{name: "Max int mod a positive equals a positive", args: args{math.MaxInt, 2}, want: 1},
		{name: "Max int mod a negative equals a negative", args: args{math.MaxInt, -2}, want: 1},
		{name: "Min int mod a positive equals a negative", args: args{math.MinInt, 3}, want: -2},
		{name: "Min int mod a negative equals a positive", args: args{math.MinInt, -3}, want: -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Modulus{}
			if got := this.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("%v %% %v = %v, wanted %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}
