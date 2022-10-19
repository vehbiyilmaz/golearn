package main

import (
	"fmt"
)

// STRUCTS are "containers" that used in golang to collect group related information (data)
// struct can hold different type of data not like array or slice
// type (ky)| name | struct keyword (ky)
// DECLARATION
type person struct {
	name   string
	age    uint16  // i did not seen negative age value till now :) 16 bit also more than enough
	kg     float32 // why use bigger float numbers it is just weight of human :)
	hasPet bool    // bool type hold as aDefault value FALSE
}

func main() {

	// INITIALIZATION
	johny := person{
		name:   "johny Desh", // declaration with ":" end of the sentece do not forget comma ","
		age:    34,
		kg:     79.8,
		hasPet: true, // even last sentence needs comma "," dont forget
	}
	// Display all content of struct
	fmt.Println(johny)
	// also single data display avaliable --> Lets display only age
	fmt.Println(johny.age) //here your Editor helps you just push . then you will se avaliable data

	// Manipulation Data
	johny.hasPet = false // johny has no more Pet :(
	fmt.Println(johny.hasPet)
}
