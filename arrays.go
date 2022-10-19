package	main

import (
"fmt"
)

func main(){

	// array is a collection of elements of the same type placed contiguos in the memory 
	// array has a fixed lenght . 

// Array DECLARATION 
//   Name   |lenght | type |elements 
var myArray = [3] int {1,2,3}

//myArray := [3] int {0,1,2} /short way 

// Different types of initializing an array 
// var myArray = [3] int {} // declared not initiallized 
// var myArray = [3] int {0,1} // partial initialization 
// var myArray = [3] int {0:0, 2:2} // index based initialization 

// Array manipulation 
myArray[0] = 1 // firt element changed tfrom 0 to  1 
myArray = [3]int {10,20,30} // all elements are changed 

fmt.Println(myArray)

}
