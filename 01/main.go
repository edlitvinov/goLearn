package main //	create package main

import (
	"fmt" //	import standart library
	"math"
)

func main() { // pointer in

	// Println automatic add in the end string "\n"
	fmt.Println("Get Start Learn Golang!")

	// Print не осуществляет перевод на новую строку
	fmt.Print("First")
	fmt.Print("Second")
	fmt.Print("Third\n")

	// declaration variables
	a := "Let's go"
	b := "learn"
	c := "Golang"

	// Printf - форматированный вывод
	fmt.Printf("%s %s %s!\n", a, b, c)

	var age int // переменная инициализирована нулем "0"
	fmt.Println("You'r age is: ", age)

	age = 37
	fmt.Println("My age: ", age)

	var fir, sec int = 13, 47
	fmt.Printf("variable Fir: %d\nvariable Sec: %d\n", fir, sec)

	// variables block
	var (
		personName     string  = "Tom"
		personLastName string  = "Greenfild"
		personAge      int     = 15
		personHeight   float64 = 1.75
		personRating   float64 = 7.5
		metric         string  = "m"
	)
	fmt.Printf("Name: %s😎\nLast Name: %s👻\nAge: %d🚴\nheight: %.2f%s🌴\nratinf: %.2f🍓\n", personName, personLastName, personAge, personHeight, metric, personRating)

	// множественное присваивание
	aArg, bArg := 15, "Bob"
	fmt.Printf("var A: %d\nvar B: %s\n", aArg, bArg)
	fmt.Printf("var A type: %T\nvar B type: %T\n", aArg, bArg)

	// add mod Math
	width, length := 20.5, 30.87
	fmt.Printf("Min decimal of rec.: %.2f\n", math.Min(width, length))

	// for example
	f, e, g, h := 13, "Red", 41.56, false
	fmt.Printf("var F: %d\nvar E: %s\nvar G: %.2f\nvar H: %t\n", f, e, g, h)
	fmt.Printf("type F: %T\ntype E: %T\ntype G: %T\ntype H: %T\n", f, e, g, h)
}
