package main

import "fmt"

// Определение простой функции
func printPrice() {
	kayakPrice := 275.00
	kayakTax := kayakPrice * 0.2
	fmt.Println("Price: ", kayakPrice)
	fmt.Println("Tax: ", kayakTax)

}
func main() {
	fmt.Println("About to coll function")
	printPrice()
	fmt.Println("Function complit")
}
