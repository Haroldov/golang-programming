/*main package - entry point*/
package main

import "fmt"

func printReverse(str string) {
	defer fmt.Print("\n")
	for _, char := range(str) {
		defer fmt.Printf("%c", char)
	}
}

func main() {
	var str string = "Hello!"

	printReverse(str)
}
