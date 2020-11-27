package main

import (
	"errors"
	"fmt"
)

func main() {
	total, err := (sum(12, 10))
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(total)
	}
}

func sum(first int, last int) (int, error) {
	if first > last {
		return 0, errors.New("First digit is greater that last digit")
	}
	total := 0
	for i := first; i <= last; i++ {
		total = total + i
	}
	return total, nil
}
