package	main

import (
"fmt"
)

func main(){

// array is a collection of elements of the same type placed contiguos in the memory 
// array has a fixed lenght. 
	
// Array DECLARATION 
//   Name   |lenght | type |elements 
var myArray = [3] int {1,2,3}
// myArray := [5] int {1,2,3,4,5} // Short way with initialization 
// myArray := [...] int {1,2,3,4,5,6} // lenght of array will be set automaticly by Complier

// Different types of initializing an array 
// var myArray = [3] int {} // declared not initiallized 
// var myArray = [3] int {0,1} // partial initialization 
// var myArray = [3] int {0:0, 2:2} // index based initialization 

// Array manipulation 
myArray[0] = 1 // firt element changed from 0 to  1 
myArray = [3]int {10,20,30} // all elements are changed 

fmt.Println("manipulated array ", myArray)

// IMPORTANT COPY ARRAYS 
// If you make new array from existing array you are making just
// new idependet copy of this array look below -->
b:= myArray
//change something usein "b" as you see Original Array is intakt 
b[0] = 23
fmt.Println("this is new array", b)
fmt.Println("this is original  array", myArray)

// POINTING Original array 
origianlArray:= [3]int{1,2,3}

// we use POINTER "&" to original Array we do not copy it we are just pointing to Original array 
copyArray := &origianlArray

//lets make some changes on COPY 
copyArray[0] = 23

// Lets see what happned to Original Array and Copy
fmt.Println("Original Array > ", origianlArray)
fmt.Println("Copy Array >", copyArray)
}
