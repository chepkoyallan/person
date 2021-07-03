package main

import (
	"fmt"
	"github.com/chepkoyallan/person/pkg/simplemath"
)

func main(){
	answer, err := simplemath.Divide(5, 3)

	if err != nil {
		fmt.Printf("An error occured %s\n", err.Error())
	} else {
		fmt.Printf("%f\n", answer)
	}

	numbers := []float64{12.2, 14, 16}
	total := simplemath.Sum(numbers...)
	sv := simplemath.NewSemanticVersion(1,2,3)
	sv.IncrementMajor()

	fmt.Printf("total of sum: %f\n", total)
	fmt.Println(sv.String())	

	// fmt.Printf("%f\n", add(2, 6));
	// fmt.Printf("%f\n", answer err);

	a := func(name string) string{
		fmt.Printf("my first %s function\n", name)
		return name
	}
	// Anonymous function
	func(){
		println("Anonymous")
	}()

	a("function ")
	a("function")

	value := a("Function 1")
	println(value)
	
	// calling returned function
	addExpr := simplemath.MathExpression("AddExpr")
	println(addExpr(2, 3))

	// Functions as parameters
	// fmt.Printf("%f", simplemath.Double(4, 6, simplemath.MathExpression("subtract")))

	// stateful functions
	p2 := simplemath.PowerOfTwo()
	value2 := p2()
	value3 := p2()
	
	println(value2, value3)

}
