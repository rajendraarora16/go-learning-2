// Assignment reader interface

package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please set a file name")
		os.Exit(1)
	}

	lw := logWriter{}

	fileName := os.Args[1]
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Something went wrong while opening file:", err)
		os.Exit(1)
	}

	io.Copy(lw, file)
}

func (logWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))

	return len(bs), nil
}
