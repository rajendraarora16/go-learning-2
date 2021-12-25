// We would like to print the value in terminal from the user...

/**
*	We would like to print a simple first and last name value...
 */

package main

import "fmt"

func main() {
	var firstName string
	var secondName string

	fmt.Println("Enter your first name!")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your second name!")
	fmt.Scanln(&secondName)

	fmt.Println("Your full name is", firstName, secondName)
}
