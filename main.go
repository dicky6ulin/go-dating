package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	//================================================================
	// variable => Sebuah penampung nilai

	// 1. variable Declaration
	var number int8 = 66
	var name string = "Dicky Agus Setiawan"
	fmt.Println(number)
	fmt.Println(name)

	var number2 = 2.66
	fmt.Printf("Bilangan Desimal = %f\n", number2)
	fmt.Printf("Bilangan Desimal = %.2f\n", number2)

	var isNumber bool = false
	fmt.Println(isNumber)

	// 2. Type inference
	nama := "Agus"
	fmt.Println(nama)

	// 3. tidak bisa di reAsign
	const jariJari = 22 / 7
	fmt.Println(jariJari)

	//================================================================
	// Scanner

	var angka int
	fmt.Print("Masukan Angka : ")
	fmt.Scan(&angka)
	fmt.Println("Angka adalah : ", angka)

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Print("Masukan Angka : ")
	// scanner.Scan()
	// input := scanner.Text()
	// inputParse, _ := strconv.Atoi(input)
	// fmt.Println("Angka adalah : ", inputParse)

	fmt.Println("===================================================")

	//Kondisi
	//for
	for i := 0; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("===================================================")
	//while

	a := 0
	for a <= 10 {
		if a%2 == 0 {
			fmt.Println(a)
		}
		a++
	}

	fmt.Println("===================================================")

	//for ever
}
