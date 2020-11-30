package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Class Class  `json:"class"`
}

type Class struct {
	Standard int `json:"class"`
	Age      int `json:"age"`
}

func main() {

	class := Class{Standard: 3, Age: 9}
	book := Student{Name: "Rithi", Class: class}

	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}
