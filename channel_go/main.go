package main

import "fmt"

func main() {

	greet := "hi there!"

	go (func(l string) {
		fmt.Println(l)
	})(greet)
}
