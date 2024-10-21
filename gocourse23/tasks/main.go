package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func incrementInt64Field(data any) {
	// Use reflect to check if data is a struct
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		fmt.Println("Error: expected a pointer to a struct")
		return
	}

	// Iterate through the fields of the struct
	val = val.Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		// Check if the field is int64
		if field.Kind() == reflect.Int64 {
			// Use unsafe to change the value of the field
			int64Ptr := (*int64)(unsafe.Pointer(val.Field(i).UnsafeAddr()))
			*int64Ptr += 1
		}
	}
}

type MyStruct struct {
	Name  string
	Count int64
}

func main() {
	myStruct := MyStruct{Name: "test", Count: 10}
	fmt.Println("Before:", myStruct)

	incrementInt64Field(&myStruct)
	fmt.Println("After:", myStruct)
}
