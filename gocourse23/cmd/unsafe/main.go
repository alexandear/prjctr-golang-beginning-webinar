package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Person struct {
	Married           bool
	Age               int32
	BankAccountAmount float64
	Name              string
	HasChildren       bool
}

func main() {
	smith := Person{Married: true, Age: 32, BankAccountAmount: 6240.5, Name: "Smith"}

	// Using Sizeof to determine the size of the structure
	fmt.Printf("Size of Person structure: %d bytes\n", unsafe.Sizeof(smith))

	// Using Offsetof to determine the offset of fields from the beginning of the structure
	fmt.Printf("Offset of field Married: %d, HasChildren: %d, Age: %d, BankAccountAmount: %d, Name: %d\n",
		unsafe.Offsetof(smith.Married),
		unsafe.Offsetof(smith.HasChildren),
		unsafe.Offsetof(smith.Age),
		unsafe.Offsetof(smith.BankAccountAmount),
		unsafe.Offsetof(smith.Name),
	)

	// Using Alignof to determine the alignment of data types
	fmt.Printf("Alignment of Married: %d, HasChildren: %d, Age: %d, BankAccountAmount: %d, Name: %d\n",
		unsafe.Alignof(smith.Married),
		unsafe.Alignof(smith.HasChildren),
		unsafe.Alignof(smith.Age),
		unsafe.Alignof(smith.BankAccountAmount),
		unsafe.Alignof(smith.Name),
	)

	// Direct access to the field using unsafe
	int32Ptr := (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&smith)) + unsafe.Offsetof(smith.Age)))
	*int32Ptr = 123
	fmt.Printf("Modified value of Age: %d\n", smith.Age)

	// Converting Pointer to *Person
	examplePtr := (*Person)(unsafe.Pointer(&smith))
	examplePtr.BankAccountAmount = 1234567890
	fmt.Printf("Modified value of BankAccountAmount: %d\n", smith.BankAccountAmount)

	bytes := []byte("Hello, world!")
	str := BytesToString(bytes)
	fmt.Println(str)
}

// BytesToString converts a slice of bytes to a string without additional memory copying.
func BytesToString(b []byte) string {
	// Getting SliceHeader, which describes the slice
	var bh = (*reflect.SliceHeader)(unsafe.Pointer(&b))

	// Creating StringHeader, which describes the string
	var sh = reflect.StringHeader{
		Data: bh.Data, // using the same data address
		Len:  bh.Len,  // the length of the slice determines the length of the string
	}

	// Converting StringHeader back to a string without copying data
	return *(*string)(unsafe.Pointer(&sh))
}
