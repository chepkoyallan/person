package simplemath

import (
	"errors"
	"math"
	// "math"
)

type MathExpr = string

const(
	AddExpr = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

// fimple function
func Multiply(p1, p2 float64) float64{
	return p1 + p2
}

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

func MathExpression(expr MathExpr) func(float64, float64) float64 {
	switch expr {
		case AddExpr:
			return Add
		case SubtractExpr:
			return Subtract
		case MultiplyExpr:
			return Multiply
		default:
			return func(f1, f2 float64) float64 {
				return 0
			}
	
	}
	// return func(f1, f2 float64) float64 {
	// 	return f1 + f2
	// }
}

// Function as paramters
func Double(f1, f2 float64, mathExpr func(float64, float64) float64) float64{
	return 2 * mathExpr(f1, f2)
}

// stateful function
func PowerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
}