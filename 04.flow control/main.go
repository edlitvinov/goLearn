package main

import "fmt"

func main() {
	var value int

	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("The nuymber", value "is even")
	}
}
