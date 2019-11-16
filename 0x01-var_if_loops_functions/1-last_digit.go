/*main package - entry point*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n, lastdigit int

	rand.Seed(time.Now().Unix())
	n = rand.Intn(10000) - 10000/2
	lastdigit = n % 10
	switch {
	case lastdigit > 5:
		fmt.Printf("Last digit of %d is %d and is greater than 5\n", n, lastdigit)
	case lastdigit == 0:
		fmt.Printf("Last digit of %d is %d and is 0\n", n, lastdigit)
	case lastdigit < 6:
		fmt.Printf("Last digit of %d is %d and is less than 6 and not 0\n", n, lastdigit)
	}
}
