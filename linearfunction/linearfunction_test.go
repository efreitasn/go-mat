package linearfunction

import (
	"strconv"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"2x+3", true},
		{"2x-3", true},
		{"2x+0", true},
		{"2x", true},
		{"-2x", true},
		{"2x + 3", true},
		{"3 + 2x", true},
		{"3 - 2x", true},
		{"3+2x", true},
		{"x", true},
		{"3", false},
		{"0x", false},
		{"2.x", false},
		{".2x", false},
		{"0x+3", false},
		{"x+3.", false},
		{"x+.3", false},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			r := IsValid(test.in)

			if r != test.out {
				t.Errorf("got %v, want %v", r, test.out)
			}
		})
	}
}

func TestNewFromString(t *testing.T) {
	tests := []struct {
		str       string
		slope     float64
		intercept float64
	}{
		{"20x+3", 20, 3},
		{"-20x+3", -20, 3},
		{"-20x-3", -20, -3},
		{"3+30x", 30, 3},
		{"-3+30x", 30, -3},
		{"-3-30x", -30, -3},
		{"3+x", 1, 3},
		{"3-x", -1, 3},
		{"-1.5-x", -1, -1.5},
		{"-1.5-2.3x", -2.3, -1.5},
		{"3x", 3, 0},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			lf := NewFromString(test.str)

			if lf.Slope() != test.slope {
				t.Errorf("got %v, want %v", lf.Slope(), test.slope)
			}

			if lf.Intercept() != test.intercept {
				t.Errorf("got %v, want %v", lf.Intercept(), test.intercept)
			}
		})
	}

	// Whether it panics when an invalid linear function string is given.
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("should've panicked")
		}
	}()

	NewFromString(".5x")
}

func TestExec(t *testing.T) {
	tests := []struct {
		str string
		in  float64
		out float64
	}{
		{"2x+4", 8, 20},
		{"x+20", 2, 22},
		{"15-2x", 5, 5},
	}

	for _, test := range tests {
		t.Run(test.str+" "+strconv.FormatFloat(test.in, 'e', 1, 64), func(t *testing.T) {
			r := NewFromString(test.str).Exec(test.in)

			if r != test.out {
				t.Errorf("got %v, want %v", r, test.out)
			}
		})
	}
}

func TestXFromY(t *testing.T) {
	tests := []struct {
		str string
		in  float64
		out float64
	}{
		{"2x+4", 8, 2},
		{"x+20", 22, 2},
		{"15-2x", 5, 5},
	}

	for _, test := range tests {
		t.Run(test.str+" "+strconv.FormatFloat(test.in, 'e', 1, 64), func(t *testing.T) {
			r := NewFromString(test.str).XFromY(test.in)

			if r != test.out {
				t.Errorf("got %v, want %v", r, test.out)
			}
		})
	}
}

func TestRoot(t *testing.T) {
	tests := []struct {
		str string
		out float64
	}{
		{"2x+4", -2},
		{"x+20", -20},
		{"15-2x", 7.5},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			r := NewFromString(test.str).Root()

			if r != test.out {
				t.Errorf("got %v, want %v", r, test.out)
			}
		})
	}
}

func TestIncreasing(t *testing.T) {
	tests := []struct {
		str string
		out bool
	}{
		{"2x+4", true},
		{"-x+20", false},
		{"15+2x", true},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			r := NewFromString(test.str).Increasing()

			if r != test.out {
				t.Errorf("got %v, want %v", r, test.out)
			}
		})
	}

}

func TestDecreasing(t *testing.T) {
	tests := []struct {
		str string
		out bool
	}{
		{"2x+4", false},
		{"-x+20", true},
		{"15+2x", false},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			r := NewFromString(test.str).Decreasing()

			if r != test.out {
				t.Errorf("got %v, want %v", r, test.out)
			}
		})
	}

}
