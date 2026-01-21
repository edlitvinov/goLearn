package main

import "fmt"

func main() {
	// Boolian
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

}
