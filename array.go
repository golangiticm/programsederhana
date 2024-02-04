package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("ini deklarasi dengan part berdasarkan alokasi indeks")
	var names [3]string
	names[0] = "Andika"
	names[1] = "Reza"
	names[2] = "monyet"

	fmt.Println(names)
	//print manual tanpa index

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//print manual dengan menyertakan index

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	//print otomatis dengan perulangan for

	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("ini deklarasi dengan part tanpa alokasi indeks")

	var fruits = [3]string{"mango", "aple", "banana"}
	fmt.Println(fruits)
	//print manual tanpa index

	fmt.Println(fruits[0])
	fmt.Println(fruits[1])
	fmt.Println(fruits[2])


	//print manual dengan menyertakan index

	for _,fruit := range fruits{
		fmt.Println(fruit)
	}
	//print otomatis dengan perulangan for
}
