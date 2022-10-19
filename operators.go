package main

import (

	"fmt"
)

func main(){

// + SUM operator operate -->int, float, complex, strings 
var a complex64 = complex(123,45)
var b complex64 = complex(123124,45345)

c := a + b 

fmt.Println(c)

// - difference like SUM operator but can not operate strings 
d:= 45
e:= 34

f := d-e
fmt.Println(c)
fmt.Println(f)
}