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

	fmt.Printf("total of sum: %f\n", total)

	// fmt.Printf("%f\n", add(2, 6));
	// fmt.Printf("%f\n", answer err);
}
