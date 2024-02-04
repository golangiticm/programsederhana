package main

import (
	"fmt"
	"strings"
)

func main() {
	days := [...]string{"sunday", "monday"}
	fmt.Println(days[0])
	fmt.Println(days[1])

	fmt.Println(strings.Repeat("=", 20))
	for i, day := range days {
		fmt.Printf("pada index ke-%d adalah hari %s\n", i, day)
	}

}
