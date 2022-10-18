package main

import "fmt"

func main() {

	// constant are unchangeable variable in go. declared with UPPERCASE
	const A int = 1 //
	const B = 3.14  // go compiler assign automaticly as float64 value

	const ( // also can be declared as a list like below

		C = 23                           // C int = 23 is also allowed
		D = 25.7                         // D float32 = 25.7 also allwed
		E = "test constant declaration " // ...

		//E = "Hello World " // E already declared and can not be changed
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
}
