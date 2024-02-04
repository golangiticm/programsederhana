package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	} //mencetak angka ber-urutan dari terkceil hingga maksimal nilai

	fmt.Println(strings.Repeat("=", 10))

	for j := 5; j > -1; j-- {
		fmt.Println("Nilai i = ", j)
	} //mencetak angka ber-urutan dari terbesar hingga minimal nilai

	fmt.Println(strings.Repeat("=", 10))

	var k, l int

	for k < 5 {
		fmt.Println("Angka ", k)
		k++
	}

	//ini perulangan menggunakan argumen kondisi
	fmt.Println(strings.Repeat("=", 13))

	for {
		fmt.Println("angka", l)
		l++

		if l == 5 {
			break
		}
	} //sedangkan ini perulangan tanpa argumen

	for i := 0; i < 5; i++ {

	}

	fmt.Println(strings.Repeat("=", 13))

	var xs = "012345678"

	for z, v := range xs {
		fmt.Println("index", z, "Value", v)
	} //pengecekan nilai ASCII dimana 0 itu = 48 byte

	for a := 0; a < 5; a++ {
		for b := a; b < 5; b++ {
			fmt.Print(b, " ")
		}
		fmt.Println()
	}

	fmt.Println(strings.Repeat("=", 13))

	for f := 0; f <= 10; f++ {
		if f%2 == 0 {
			fmt.Println("Angka : ", f)
		}
	} // mencetak angka genap


	
	fmt.Println(strings.Repeat("=", 13))

	for g := 0; g <= 10; g++ {
		if g%2 == 1 {
			fmt.Println("Angka : ", g)
		}
	} // mencetak angka ganjil


	
	fmt.Println(strings.Repeat("=", 13))

	for g := 10; g >0; g-- {
		if g%2 == 1 {
			fmt.Println("Angka : ", g)
		}
	} // mencetak angka ganjil
}
