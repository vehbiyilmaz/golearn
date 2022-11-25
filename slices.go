package main

import (
	"fmt"
)

func main (){

// Slices works like Arrays, but more flexible 
// Do not forget every slice is a part of an array 
// even you create a directly slice in Go. in background Go creates an array then take 
// part of this array as a Slice 

// Slice Declaration 
var mySlice = []int{1,2,3,4,5,6,7} // Commun way declaration and initialization 

// mySlice := []int{1,2,3,4,5} // shorter way 
// mySlice := make([]int,5,5)  // creation with make func. arguments ->length , ->capacity

// Create a Slice from an Array 
var myArray = [5] int{1,2,3,4,5} // array with 5 element 
mySli := myArray[1:3] // create slice from array  intex 1 till index3  <- " thid not included"

// Manupulation Slices 
mySlice[1] =20  // changing index1 from 2 to 20 

// Add new element to Slide "APPEND"
// Do not forget before we append some element to slice Capacity of Slice were 7
// after adding just 1 element Slice length will be  8  
// but capacity  will be double of original Slice so 7x2 -> 14 

// cap() function to see capacity of Slice 
// also len() function avaliable to see length of slice 

fmt.Println("Before adding element cap. is = ",cap(mySlice))
mySlice = append(mySlice, 8)
fmt.Println("after adding element Slice looks like ", mySlice)
fmt.Println("after adding element to slice now capacity is : ",cap(mySlice))

// Also we can create a slice from existing 2 o more slices just grouping them 
sliceOne := []int {1,2,3}
sliceTwo := []int {4,5,6}
newSlice:= append(sliceOne,sliceTwo...) // new Slice groups 2 Slices 

// IMPORTANT  SLICE COPY IS A JUST POINT ORIGINAL SLICE WITH NEW NAME 
// when you make a slice copy  does mean you point Original slice
// every change on new created slice copy will qaffected to original slice  
newSlice[0] = 23

fmt.Println(mySlice)
fmt.Println("Slice created from Array ",mySli)
fmt.Println(newSlice)

}