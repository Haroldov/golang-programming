/*main package - entry point*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int

	rand.Seed(time.Now().Unix())
	n = rand.Intn(10000) - 10000/2
	switch {
	case n == 0:
		fmt.Println(n, "is zero")
	case n > 0:
		fmt.Println(n, "is positive")
	case n < 0:
		fmt.Println(n, "is negative")
	}
}
