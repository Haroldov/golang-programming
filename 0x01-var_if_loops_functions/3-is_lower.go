/*main package - entry point*/
package main

import "fmt"

func isLower(n int) bool {
	if n >= 97 && n < 123 {
		return true
	}
	return false
}

func main() {
	var char int = 100

	res := isLower(char)
	fmt.Println(res)
}
