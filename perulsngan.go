package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	} //mencetak angka ber-urutan dari terkceil hingga maksimal nilai

fmt.Println(strings.Repeat("=",10))

	for j := 5; j >-1; j-- {
		fmt.Println("Nilai i = ", j)
	} //mencetak angka ber-urutan dari terbesar hingga minimal nilai

fmt.Println(strings.Repeat("=",10))
  

var k, l int

for k < 5{
	fmt.Println("Angka ", k)
	k++
}

 //ini perulangan menggunakan argumen kondisi
fmt.Println(strings.Repeat("=",13))

for {
	fmt.Println("angka", l)
	l++

	if l == 5 {
		break
	}
} //sedangkan ini perulangan tanpa argumen 

}
