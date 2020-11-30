package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile() {

	fmt.Printf("Creating a file")

	file, err := os.Create("mythily.txt")

	if err != nil {
		log.Fatalf("failed creating a file: %s", err)
	}

	defer file.Close()

	len, err := file.WriteString("Hello")

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func ReadFile() {

	fmt.Printf("\n\nReading a file")
	fileName := "test.txt"

	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}

func main() {

	CreateFile()
	ReadFile()
}
