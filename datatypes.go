package main

import "fmt"

func main(){

// INTEGER not fractional nuimbers, positive or negative value 
var a int = 8 // Size depend on Platform (32/64bit) 
var b int8 = 127 // between -128 and 127
var c int16 = 32767 // between -32768 and 32767
var d int32 = 234233245 // 32 bit 
var e int64 = 3425324532453543 // 64 bit 

// UNIT Unsigned INTEGER ONLY POSITIVE numbers 
var f uint = 2342 // // Size depend on Platform (32/64bit)
var g uint8 = 255 // Only positiv numbers 8 bit 
var h uint16 =  65535 // positiv numbers till 65535 16 bit 
var g uint32 = 63482764 // 32 bit positive numbers 
var h uint64 = 26387687687 // 64 bit positive numbers 
var i uintptr = 56576//special unsigned intiger LARGE ENOUGHT to hold any POINTER ADRESS platform dependent 32/64

// FLOATS fraccional numbers positive or negative 
var j float32 = 13231241341.545353453453 // 32 bit 
var k float64 = 873264827634872634872.234324235 // 64 bit 

// STRINGS sequence of characters 
var l string = "test12" // declared between ""

// Complex Numbers --> combination of real and imaginary numbers 
// first half real and second half imaginary part 
var m complex64 = complex(17242,-2398)
var n complex128 = complex(723453,-245348)


// BYTE is unsignet intiger of 8 bits only POSITIVE numbers 
var o byte = 255 // numbers betwwen o and 255 included like uint8

// RUNE just like int32 bit 
var p rune = -98799879

}


