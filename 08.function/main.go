package main

import "fmt"

// Определение простой функции
// func printPrice() {
// 	kayakPrice := 275.00
// 	kayakTax := kayakPrice * 0.2
// 	fmt.Println("Price: ", kayakPrice)
// 	fmt.Println("Tax: ", kayakTax)

// Определение и использование параметров функции
func printPrice(product string, price, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Printf("%s\nprice: %.2f\ntax: %.2f\n", product, price, taxAmount)
	fmt.Println("---------------")
}

func main() {
	// Вызов функции printPrice()
	// fmt.Println("About to coll function")
	// printPrice()
	// fmt.Println("Function complit")

	// Вызов функции printPrice() с аргументами

	printPrice("Kayak", 350, 0.5)
	printPrice("LifeJacket", 37.85, 0.3)
	printPrice("Soccer Ball", 17.50, 0.1)

}
