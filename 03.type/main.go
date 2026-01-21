package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Boolian
	fmt.Println("*Type Boolian*")
	fmt.Println("--------------")
	var firstBool bool // default false

	fmt.Printf("variable \"firstBool\" is value: %t\n", firstBool)
	fmt.Printf("variable \"firstBool\" is type: %T\n", firstBool)

	aBool, bBool := true, false
	fmt.Printf("variable \"aBool\" is value: %t\n", aBool)
	fmt.Printf("variable \"bBool\" is value: %t\n", bBool)
	fmt.Printf("aBool && (AND) bBool is: %t\n", aBool && bBool)
	fmt.Printf("aBool || (OR) bBool is: %t\n", aBool || bBool)
	fmt.Printf("!(NOT) aBool is: %t\n", !aBool)
	fmt.Printf("!(NOT) bBool is: %t\n", !bBool)
	fmt.Println("------------------------------------------")
	// Numerics
	// Int
	fmt.Println("*Type Integer*")
	fmt.Println("--------------")
	fmt.Printf("Integers type (signed): \n- int\n- int8\n- int16\n- int32\n- int64\n")
	fmt.Printf("Integers type (unsigned): \n- uint\n- uint8\n- uint16\n- uint32\n- uint64\n")

	var a int
	var b int8
	var c int16
	var d int32
	var e int64

	fmt.Printf("Sizeof variable \"int\" type is: %d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("Sizeof variable \"int8\" type is: %d bytes\n", unsafe.Sizeof(b))
	fmt.Printf("Sizeof variable \"int16\" type is: %d bytes\n", unsafe.Sizeof(c))
	fmt.Printf("Sizeof variable \"int32\" type is: %d bytes\n", unsafe.Sizeof(d))
	fmt.Printf("Sizeof variable \"int64\" type is: %d bytes\n", unsafe.Sizeof(e))
	fmt.Println("------------------------------------------")
	fmt.Printf("Sizeof int == int64 == 8 bytes\n")
	fmt.Printf("But type int !== type int64\n")
	fmt.Println("Sizeof signed int == sizeof unsigned int")
	fmt.Println("------------------------------------------")

	// Numeric. Float
	fmt.Println("*Float Type*")
	fmt.Println("------------")
	fmt.Printf("Float Type in Golang is:\n- float32\n- float64\n")

}
