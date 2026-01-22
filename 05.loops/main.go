package main

import "fmt"

func main() {
	// loop FOR
	fmt.Println("*Loop FOR*")
	fmt.Println("------------")
	for i := 0; i <= 5; i++ {
		fmt.Printf("i : %d\n", i)
	}
	fmt.Println()
	fmt.Println("used BREAK")
	fmt.Println("------------")
	for i := 0; i < 15; i++ {
		if i > 7 {
			break
		}
		fmt.Printf("i : %d\n", i)
	}

	fmt.Println("After for loop with break")
	fmt.Println()
	fmt.Println("used Continue")
	fmt.Println("--------------")

	for i := 0; i < 12; i++ {

		if i == 4 || i == 7 || i == 9 {
			continue
		}
		fmt.Printf("i : %d\n", i)
	}

	fmt.Println("After for loop with continue")
	fmt.Println()
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("*triangle*")
	fmt.Println()
	// while loop
	fmt.Println("*While loop*")
	fmt.Println("-------------")
	fmt.Print()
	var loopVar int = 0
	for loopVar < 7 {
		fmt.Printf("while loop: %d\n", loopVar)
		loopVar++
	}
	fmt.Println()
	for x, y := 0, 1; x <= 10 && y <= 12; x, y = x+1, y+2 {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}

}
