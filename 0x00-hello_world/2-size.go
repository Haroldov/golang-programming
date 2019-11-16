/*main package - entry point*/
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("Size of an int: %d byte(s)\n", unsafe.Sizeof(int(5)))
}
