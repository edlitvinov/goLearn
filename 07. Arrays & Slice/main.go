package main

import "fmt"

func main() {
	fmt.Println("*Arrays*")
	fmt.Println("---------")
	var firstArr [5]int // создание массива из 5 int-ов
	fmt.Println("Array: ", firstArr)
	fmt.Println()
	var newArr = [7]int{13, 47, -5, 89, -52, 7, 35}
	fmt.Println("New Array: ", newArr)
	fmt.Println()
	addArr := [...]int{14, 57, 23, -6, 12}
	fmt.Println("Array [...]: ", addArr)
	lArr := len(addArr)
	fmt.Println("Len Array: ", lArr)
	fmt.Println()
	fArr := [...]float64{1.54, 15.75, 65.02, 7.89, 8.54, 9.7}
	for i := 0; i < len(fArr); i++ {
		fmt.Printf("%d element array: %.2f\n", i, fArr[i])
	}
	fmt.Println()
	var sum float64
	for id, val := range fArr {
		fmt.Printf("%d element Array: %.2f\n", id, val)
		sum += val
	}
	fmt.Println("----------------------")
	fmt.Printf("          sum: %.2f\n", sum)
	fmt.Println()
	for _, val := range fArr {
		fmt.Printf(" %.2f\n", val)
	}
	fmt.Println()
	wordsName := [2][2]string{
		{"Tom", "John"},
		{"Boby", "Alisa"},
	}
	fmt.Println("mega Array: ", wordsName)
	fmt.Println()
	for _, val1 := range wordsName {
		for _, val2 := range val1 {
			fmt.Printf("%s\n", val2)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("*Slices*")
	fmt.Println("---------")
	// Slise - это указатель на массив
	fmt.Println()
	slice := []int{1, 23, 45, 74}
	fmt.Printf("slice: %d\n", slice)
	slice = append(slice, 98)
	fmt.Printf("slice: %d\n", slice)
	fmt.Println()
	startArr := [4]int{10, 20, 30, 40}
	var startSlice []int = startArr[0:2]
	fmt.Println("slice [0:2] :", startSlice)
	fmt.Println()
	s1 := make([]int, 7, 14)
	fmt.Println(s1)

}
