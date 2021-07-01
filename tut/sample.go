package tut

import (
	// "errors"
	"fmt"

	// "github.com/chepkoyallan/person/pkg/create/person"
)

func main(){
	//Invoking the function
	// allan := createPerson("Allan", "Kiplangat", "Chepkoy")
	fmt.Println()
}

func loopTo() (error){
	var i int
	//loop to
	for i < 5 {
		i++	
		//break in the loop early
		if i == 3 {
			break
			//continue
		}
	}
	return nil;
}

func loopToWithPostClause() (error){
	//var i int
	for i := 0; i < 5; i++ {
		println(i)
	}
	return nil;
}

func infiniteLoop() (error){
	//var i int
	for i := 0; i < 5; i++ {
		println(i)
	}
	return nil;
}

func forLoop() (error){
	var i int
	for {
		if i == 5 {
			break
		}
		print(i)
		i++
	}
	return nil;
}

// looping through a collection (Works with all collections)
func loopCollection(){
	myList := []int{1, 2, 3}
	for i, v := range myList {
		println(i, v)
	}
}

func ifStatement(){
	myList := []int{1, 2, 3}
	for i, v := range myList {
		println(i, v)
	}
	// equal
	if 0 == 0 {
		println("Its the same")
	}
	// Not equal
	if 0 != 0 {
		println("Its the same")
	}
	// Not equal
	if 0 != 0 {
		println("Its the same")
	}
	// Not equal else
	if 0 != 0 {
		println("Its the same")
	} else {
		println("Different users")
	}


}

func switchStatement(){
	var i string

	switch i {
		case "GET":
			println("")
			fallthrough
		case "POST":
			println("")
		default:
			println("")
	}
}