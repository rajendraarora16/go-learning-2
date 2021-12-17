package main

import "fmt"

type wordArr []string

func main() {

	sliceArr := wordArr{"hi", "my", "name", "is", "Raj"}

	updateValue(sliceArr)
	sliceArr.printVal()
}

func updateValue(val []string) {
	val[0] = "Bye"
}

func (w wordArr) printVal() {
	fmt.Println(w)
}
