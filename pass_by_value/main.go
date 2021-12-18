package main
import "fmt"

func main() {
 n := "Raj"
 nPointer := &n
 fmt.Println(&nPointer)
 printPointer(nPointer)
}
 
func printPointer(nPointer *string) {
 	fmt.Println(&nPointer)
}