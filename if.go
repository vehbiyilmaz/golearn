package main

import (
	"fmt"
)

func main(){

	if 2*2 == 4 {
		fmt.Println("it is true 2x2 is 4 ")
	}

	if myNum := 6 ; myNum > 0 {

		fmt.Println("Your number ist a positive number")
		
	} else if myNum == 0 {
		fmt.Println(" Your number ist ZERO")

	} else {
		fmt.Println("Your number ist NEGATIVE")
	}
}