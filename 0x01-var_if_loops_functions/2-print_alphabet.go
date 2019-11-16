/*main package - entry point*/
package main

import "fmt"

func main() {
	var char int = 97

	for ; char < 123; char++ {
		fmt.Printf("%c", char)
	}
	fmt.Print("\n")
}
