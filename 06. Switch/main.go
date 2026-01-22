package main

import "fmt"

func main() {
	fmt.Println("*Switch*")
	fmt.Println("--------")
	fmt.Println()

	var price int
	fmt.Println("input price: (100, 125, 130, 132, any number)")
	fmt.Scan(&price)

	switch price {
	case 100:
		fmt.Println("varic #1")
	case 125:
		fmt.Println("varic #2")
	case 130:
		fmt.Println("varic #3")
	case 132:
		fmt.Println("varic #4")
	default:
		fmt.Println("default varic")
	}

	fmt.Println()
	var ageGroup string = "q"
	switch ageGroup {
	case "a", "b", "c":
		fmt.Println("Age Group 30-40")
	case "d", "f", "g":
		fmt.Println("Age Group 18 - 28")
	case "e", "h", "i":
		fmt.Println("Age Group 45-50")
	default:
		fmt.Println("default group")
	}

	fmt.Println()
	fmt.Println("Сколько вам лет?: ")
	var age int
	fmt.Scan(&age)

	switch {
	case age > 10 && age <= 18:
		fmt.Println("You youg!")
	case age > 18 && age <= 25:
		fmt.Println("Ты то что нам нужно!")
	case age > 25 && age <= 35:
		fmt.Println("Ты уже опытный крендель!")
	case age > 35 && age <= 50:
		fmt.Println("Ты зрелый перчик!")
	case age > 50 && age <= 60:
		fmt.Println("Ты учитель!")
	default:
		fmt.Println("Твой возраст самый лучший!")
	}
	fmt.Println()
	var num int
	fmt.Println("Ведите любое сило до 100:")
	fmt.Scan(&num)

	switch {
	case num < 50:
		fmt.Printf("Вы ввели число: %d\n", num)
		fallthrough // "проваливание" осуществляется без проверки условий
	case num < 100:
		fmt.Printf("Вы ввели число: %d\n", num)
		break // принудительная остановка проваливания
	}
	fmt.Println()
	var i int
uberloop: // label использовать только во вложенные циклы
	for {
		fmt.Println("Введите любое число: ")
		fmt.Scan(&i)
		switch {
		case i%2 == 0:
			fmt.Printf("Вы ввели %d\n%d - четное число\n", i, i)
			break uberloop
		}
	}
	fmt.Println("THE END !")

}
