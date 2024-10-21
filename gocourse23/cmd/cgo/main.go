package main

/*
#include <stdio.h>
#include <stdlib.h>

#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lmathlib
#include "mathlib.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	a := 10
	b := 5

	sum := C.add(C.int(a), C.int(b))
	quotient := C.divide(C.int(a), C.int(b))
	difference := C.subtract(C.int(a), C.int(b))
	product := C.multiply(C.int(a), C.int(b))

	fmt.Printf("Sum for %T: %d\n", sum, sum)
	fmt.Printf("Quotient for %T: %f\n", quotient, quotient)
	fmt.Printf("Difference: %d\n", difference)
	fmt.Printf("Product: %d\n", product)

	for {
		s := `Some danger string`
		cString := C.CString(s)
		fmt.Println(cString)
		C.free(unsafe.Pointer(cString))
	}
}
