// Package linearfunction provides functions to work with linear functions.
package linearfunction

import (
	"regexp"
	"strconv"
	"strings"
)

// RegExps
var linearFnStrRx = regexp.MustCompile(`^(?:(?:(?:\+|-)?(?:(?:(?:[1-9]{1,}[0-9]*)?x)|(?:[0-9]{1,}\.[0-9]{1,}x))(?:(?:\+|-)[0-9]{1,}[0-9]*(?:\.[0-9]{1,})?)?)|(?:(?:\+|-)?[0-9]{1,}[0-9]*(?:\.[0-9]{1,})?(?:\+|-)(?:(?:(?:[1-9]{1,}[0-9]*)?x)|(?:[0-9]{1,}\.[0-9]{1,}x))))$`)
var slopeRx = regexp.MustCompile(`(\+|-)?(?:(?:([1-9]{1,}[0-9]*)?x)|(?:([0-9]{1,}\.[0-9]{1,})x))`)
var interceptRx = regexp.MustCompile(`(\+|-)?([0-9]{1,}[0-9]*(?:\.[0-9]{1,})?)`)

// LinearFunction is a linear function.
type LinearFunction struct {
	slope     float64
	intercept float64
}

// IsValid returns whether a string is a valid linear function.
func IsValid(str string) bool {
	return linearFnStrRx.MatchString(removeSpaces(str))
}

func removeSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

// NewFromString returns a linear function from a string.
func NewFromString(str string) LinearFunction {
	if !IsValid(str) {
		panic("linearfunction: invalid linear function string")
	}

	str = removeSpaces(str)
	slopeMatches := slopeRx.FindStringSubmatch(str)
	slope := float64(1)

	if slopeMatches[2] != "" {
		slope, _ = strconv.ParseFloat(slopeMatches[2], 64)
	} else if slopeMatches[3] != "" {
		slope, _ = strconv.ParseFloat(slopeMatches[3], 64)
	}

	if slopeMatches[1] == "-" {
		slope *= -1
	}

	strWithoutSlope := strings.ReplaceAll(str, slopeMatches[0], "")
	interceptMatches := interceptRx.FindStringSubmatch(strWithoutSlope)
	intercept := float64(0)

	if len(interceptMatches) > 0 {
		intercept, _ = strconv.ParseFloat(interceptMatches[2], 64)

		if interceptMatches[1] == "-" {
			intercept *= -1
		}
	}

	return LinearFunction{
		slope:     slope,
		intercept: intercept,
	}
}

// Slope returns the slope value of the linear function.
func (lf LinearFunction) Slope() float64 {
	return lf.slope
}

// Intercept returns the intercept value of the linear function.
func (lf LinearFunction) Intercept() float64 {
	return lf.intercept
}

// Exec executes the linear function with the provided input.
func (lf LinearFunction) Exec(x float64) float64 {
	return (lf.Slope() * x) + lf.Intercept()
}

// XFromY returns the value of x when y equals the provided value.
func (lf LinearFunction) XFromY(y float64) float64 {
	return (y - lf.Intercept()) / lf.Slope()
}

// Root returns the value of x when y equals 0.
func (lf LinearFunction) Root() float64 {
	return lf.XFromY(0)
}

// Increasing returns whether the linear function is increasing.
func (lf LinearFunction) Increasing() bool {
	return lf.Slope() > 0
}

// Decreasing returns whether the linear function is decreasing.
func (lf LinearFunction) Decreasing() bool {
	return !lf.Increasing()
}
