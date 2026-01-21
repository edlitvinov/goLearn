package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
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

	// Strings
	fmt.Println("*Type Strings*")
	fmt.Println("------------")
	name := "Tom"
	nameRu := "Том"
	lastName := "Sawyer"
	fullName := name + " " + lastName
	fmt.Println("My name:", fullName)
	fmt.Println("Length of string:", name, len(name))
	fmt.Println("Length of string:", nameRu, len(nameRu))
	fmt.Println("Rune - 1  \"utf\" symbol")
	fmt.Println("*Used Rune*")
	fmt.Println("Length of string:", name, len(name))
	fmt.Println("Length of string:", nameRu, utf8.RuneCountInString(nameRu))
	fmt.Println("--------------------")
	totalString, subString := "Tom Sawyer", "om"
	fmt.Println(strings.Contains(totalString, subString))
	fmt.Println("Rune -> alias on type INT32")
	var firstRune rune
	var secondRune rune = 'F'
	var thirdRune rune = 47
	fmt.Printf("value variable firstRune: %c\n", firstRune)
	fmt.Printf("value variable secondRune: %c\n", secondRune)
	fmt.Printf("value variable thirdRune: %c\n", thirdRune)
	fmt.Printf("Sizeof Rune: %d\n", unsafe.Sizeof(firstRune))
	fmt.Println("Type BYTE -> alias type UINT8")
}
