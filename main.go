package main //	create package

import "fmt" //	import standart library

func main() { // pointer in

	// Println automatic the end string "\n"
	fmt.Println("Get Start Learn Golang!")

	// Print не осуществляет перевод на новую строку
	fmt.Print("Fisst")
	fmt.Print("Second")
	fmt.Print("Third\n")

	// Printf - форматированный вывод
	a := "Let's go"
	b := "learn"
	c := "Golang"

	fmt.Printf("%s %s %s!\n", a, b, c)

}
