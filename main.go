package main

import (
	"fmt"
)

func main() {

	//Defer

	// defer fmt.Println("start")
	// defer fmt.Println("1")
	// defer fmt.Println("2")
	// fmt.Println("Proses")

	// Proses
	// 2
	// 1
	// start

	//=================================================================
	// ARRAY

	var numbers [3]int

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3

	for i, v := range numbers {
		fmt.Printf("Index : %d , Value : %d\n", i, v)
	}

	fmt.Println("===========")

	numbers1 := [3]int{1, 2, 3}
	for i, v := range numbers1 {
		fmt.Printf("Index : %d , Value : %d\n", i, v)
	}

	fmt.Println("===========")

	var numbers2 = [3]int{1, 2, 3}
	for i, v := range numbers2 {
		fmt.Printf("Index : %d , Value : %d\n", i, v)
	}

	fmt.Println("=========== Variadic ==========")

	var numbersVariadic = [...]int{1, 2, 3}
	for _, v := range numbersVariadic {
		fmt.Printf("Value : %d\n", v)
	}

	fmt.Println("=========== Array Multidimensi ==========")

	var numbersAM = [2][5]int{{1, 2, 3, 4, 5}, {6, 7, 8}}
	for _, row := range numbersAM {
		for _, numberSatuan := range row {
			fmt.Printf("Value : %d\n", numberSatuan)
		}
	}

}
