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

// указатели в качестве параметров
func swap(first, second *int) {
	fmt.Printf("Before swap:\nfirst: %v\nsecond: %v\n", *first, *second)
	temp := *first
	*first = *second
	*second = temp
	fmt.Printf("After swap:\nfirst: %v\nsecond: %v\n", *first, *second)
}

// Использование результатов функции
func calcTax(price float64) float64 {
	return price + (price * 0.2)
}

// возврат функцией нескольких результатов
func swapVal(first, second int) (int, int) {
	return second, first
}

func main() {
	// Вызов функции printPrice()
	fmt.Println("About to coll function")
	price()
	fmt.Println("Function complit")
	fmt.Println("------------------")
	fmt.Println("")
	// Вызов функции printPrice() с аргументами
	fmt.Println("Defining and Using Function Parameters")
	fmt.Println("---------------------------------------")
	printPrice("Kayak", 350, 0.5)
	printPrice("LifeJacket", 37.85, 0.3)
	printPrice("Soccer Ball", 17.50, 0.1)
	fmt.Println("")
	//вызов функции с пустым идентификатором
	fmt.Println("Omitting Parameter Names")
	fmt.Println("-------------------------")
	newProd("Kayak", 380, 0.2)
	newProd("LifeJacket", 37.45, 0.1)
	newProd("Kayak", 20.40, 0.15)
	fmt.Println("------------------")
	fmt.Println("")
	// вызов функции с вариативными параметрами, вместо массива можно применить "..."
	fmt.Println("Dealing with No Arg's for a Variadic Param.")
	fmt.Println("--------------------------------------------")
	nameSupp := []string{"Acme Kayaks", "Bob's Boats", "Crayzy Canors"}
	//printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crayzy Canors")
	printSuppliers("Kayak", nameSupp...)
	printSuppliers("LifeJacket", "Sail Safe Co")
	printSuppliers("Sail Safe Co")
	fmt.Println("")
	fmt.Println("Указатели в качестве параметров")
	fmt.Println("--------------------------------")
	num1, num2 := 12, 17
	fmt.Printf("value NUM1: %d\nvalue NUM2: %d\n", num1, num2)
	swap(&num1, &num2)
	fmt.Printf("After SWAP:\nNUM1: %d\nNUM2: %d\n", num1, num2)
	fmt.Println("")
	fmt.Println("---------------------")
	products := map[string]float64{
		"Kayak":      275.45,
		"LifeJacket": 35.85,
	}

	for product, price := range products {
		priceWithTax := calcTax(price)
		fmt.Printf("Product: %s\nPrice: %.2f\n", product, priceWithTax)
	}
	fmt.Println("---------------------")
	fmt.Println("")
	fmt.Println("Return Multiple Function Results")
	fmt.Println("---------------------------------")
	v1, v2 := 37, 48
	fmt.Printf("Before calling func:\nv1: %d\nv2: %d\n", v1, v2)
	v1, v2 = swapVal(v1, v2)
	fmt.Printf("After calling func:\nv1: %d\nv2: %d\n", v1, v2)

}
