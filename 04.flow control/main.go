package main

import "fmt"

func main() {
	var value int

	fmt.Println("input any number")
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("The nuymber", value, "is even")
	} else {
		fmt.Println("The nuymber", value, "is odd")
	}

	// else-if
	// if cond1 {
	// 	...
	// } else if cond2 {
	// 	...
	// } else if cond3 {
	// 	...
	// } else if cond4 {
	// 	...
	// } else if ... {
	// 	...
	// } else {
	// 	...
	// }
	if num := 15; num%2 == 0 {
		fmt.Println(num)
		fmt.Println("EVEN")
	} else {
		fmt.Println(num)
		fmt.Println("ODD")
	}

}
