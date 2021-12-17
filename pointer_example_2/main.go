package main

import "fmt"

func main() {
	w := "raj"

	updateWord(w)

	fmt.Println(w)
}

func updateWord(w string) {
	w = "Bill"
}
