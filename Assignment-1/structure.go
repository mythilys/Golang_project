--structure in go

package main

import (
	"fmt"
)

type chocolate struct {
	firstName string
	lastName  string
	price     int
	quantity  int
}

func main() {
	cho1 := chocolate{
		firstName: "Kit",
		lastName:  "Kat",
		price:     25,
		quantity:  10,
	}
	fmt.Println("chocolate", cho1)
}
