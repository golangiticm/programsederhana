package main

import "fmt"



func main() {
	var umur int
	var pilihan int



	fmt.Println("Masukkan pilihan Anda ")
	fmt.Println("1. Cek Group umur\n2. Exit")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
	fmt.Print("Masukkan Umur Anda : ")
	fmt.Scanln(&umur)

	if umur  <= 10 {
		fmt.Println("Anda Masih kanank-kanak!")
	} else if umur <= 20 {
		fmt.Println("Anda Remaja")
	}else if umur <= 30 {
		fmt.Println("Anda Dewasa")
	}else if umur > 30 {
		fmt.Println("Anda sudah lansia")
	}

    default:
		break
	}

}
//ouputannya ini 
//umur anda sekian anda masuk dalam kategori dewqas misal 