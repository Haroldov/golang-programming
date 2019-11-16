/*main package - entry point*/
package main

import "fmt"

func is_lower(n int) bool {
	if n >= 97 && n < 123 {
		return true
	} else {
		return false
	}
}



func main() {
	var char int = 100

	res := is_lower(char)
	fmt.Println(res)
}
