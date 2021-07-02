package simplemath

import (
	"errors"
	// "math"
)

// fimple function
func Add(p1, p2 float64) float64{
	return p1 + p2
}

// fimple function
func Subtract(p1, p2 float64) float64{
	return p1 - p2
}

//return multiple values from the function
func Divide(p1, p2 float64) (answer float64, err error) {
	// short circuiting
	if p2 == 0 {
		println("Hello")
		err = errors.New("operation devided by zero not allowed")
	}
	answer = p1 / p2
	return
}

// Variadic function (need to be the last parameters)
func Sum(values ...float64) float64{
	total := 0.0
	for _, value := range values{
		total += value
	}
	return total
}