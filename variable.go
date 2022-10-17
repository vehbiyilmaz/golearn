package main

import (
	"fmt"
)

func main() {

	// 3 ways to declare variable a -b -c
	var a int // declaration  default value is ZERO if it's string "" empty string
	a = 64    // initialization
	fmt.Println("a is : ", a)

	var b int = 8 // declaration and initialization
	fmt.Println("b is : ", b)

	var c = 8.8 // golang compiler recognize as a float and assign it as a float64
	fmt.Println("c is : ", c)

	//Data convertion
	d := int(c) // dont forget convertig float number to integer you lose some data
	fmt.Println("d is  value of c but INT : ", d)

}
