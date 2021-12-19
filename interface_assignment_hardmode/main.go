// Assignment reader interface

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please set a file name")
		os.Exit(1)
	}

	fileName := os.Args[1]
	fileContents, err := OpenFile(fileName)

	if err != nil {
		fmt.Println("Something went wrong while opening file:", err)
	}

	buf := make([]byte, 99999)
	fileContents.Read(buf)
	fmt.Println(string(buf))
}

func OpenFile(fileName string) (*os.File, error) {

	file, err := os.Open(fileName)
	return file, err
}
