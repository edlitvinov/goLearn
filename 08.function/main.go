package main

import "fmt"

// Определение простой функции
func price() {
	kayakPrice := 275.00
	kayakTax := kayakPrice * 0.2
	fmt.Println("Price: ", kayakPrice)
	fmt.Println("Tax: ", kayakTax)
}

// Определение и использование параметров функции
func printPrice(product string, price, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Printf("%s\nprice: %.2f\ntax: %.2f\n", product, price, taxAmount)
	fmt.Println("---------------")
}

// пропуск имен параметров
// пустой идентификатор _
func newProd(product string, price, _ float64) {
	taxAmount := price * 0.15
	fmt.Printf("%s\nprace: %.2f\nTax: %.2f\n", product, price, taxAmount)
}

// функция с вариативными параметрами
func printSuppliers(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Printf("Product: %s\nSuppliers: none!", product)
	} else {
		for _, supplier := range suppliers {
			fmt.Printf("Product: %s\nSupplier: %s\n", product, supplier)
		}
	}
}

func main() {
	// Вызов функции printPrice()
	fmt.Println("About to coll function")
	price()
	fmt.Println("Function complit")
	fmt.Println("------------------")

	// Вызов функции printPrice() с аргументами
	fmt.Println("Defining and Using Function Parameters")
	fmt.Println("---------------------------------------")
	printPrice("Kayak", 350, 0.5)
	printPrice("LifeJacket", 37.85, 0.3)
	printPrice("Soccer Ball", 17.50, 0.1)

	//вызов функции с пустым идентификатором
	fmt.Println("Omitting Parameter Names")
	fmt.Println("-------------------------")
	newProd("Kayak", 380, 0.2)
	newProd("LifeJacket", 37.45, 0.1)
	newProd("Kayak", 20.40, 0.15)
	fmt.Println("------------------")

	// вызов функции с вариативными параметрами, вместо массива можно применить "..."
	fmt.Println("Dealing with No Arg's for a Variadic Param.")
	fmt.Println("--------------------------------------------")
	printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crayzy Canors")
	printSuppliers("LifeJacket", "Sail Safe Co")
	printSuppliers("Sail Safe Co")

}
