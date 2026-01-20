package main

import "fmt"

func main() {
	var (
		age  int
		name string
	)

	fmt.Printf("What is you name?\n")
	fmt.Scan(&name)
	fmt.Printf("How old are you?\n")
	fmt.Scan(&age)

	fmt.Printf("You'r Name is: %s\nYou'r Age is: %d", name, age)

}
