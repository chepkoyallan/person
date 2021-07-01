package main

import (
	"fmt"

	"github.com/chepkoyallan/person/pkg/organization"
)

func main(){
	// p := organization.NewPerson("Allan", "Chepkoy", organization.NewSocialSecurityNumber("123-45-6789"))
	p := organization.NewPerson("Allan", "Chepkoy", organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany"))
	err:=p.SetTwitterHandler("@allan_chepkoy")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occured setting twitter handler: %s\n", err.Error())
	}
	// println(p.TwitterHandler())
	// println(p.TwitterHandler().RedirectUrl())
	// println(p.ID())
	// println(p.FullName())
	// println(p.Name.FullName())
	// println(p.Country())

	//Equality


	
}
