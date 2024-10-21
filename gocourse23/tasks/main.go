package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func incrementInt64Field(data any) {
	// Використовуємо reflect для перевірки, чи є data структурою
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		fmt.Println("Помилка: очікується покажчик на структуру")
		return
	}

	// Ітеруємо через поля структури
	val = val.Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		// Перевіряємо, чи поле є int64
		if field.Kind() == reflect.Int64 {
			// Використовуємо unsafe для зміни значення поля
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
	fmt.Println("До: ", myStruct)

	incrementInt64Field(&myStruct)
	fmt.Println("Після: ", myStruct)
}
