package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Person struct {
	Maried            bool
	Age               int32
	BancAccountAmount float64
	Name              string
	HasChildren       bool
}

func main() {
	Smith := Person{Maried: true, Age: 32, BancAccountAmount: 6240.5, Name: "Smith"}

	// Використання Sizeof для визначення розміру структури
	fmt.Printf("Розмір структури Person: %d байтів\n", unsafe.Sizeof(Smith))

	// Використання Offsetof для визначення зміщення полів від початку структури
	fmt.Printf("Зміщення поля Maried: %d, HasChildren: %d, Age: %d, BancAccountAmount: %d, Name: %d\n",
		unsafe.Offsetof(Smith.Maried),
		unsafe.Offsetof(Smith.HasChildren),
		unsafe.Offsetof(Smith.Age),
		unsafe.Offsetof(Smith.BancAccountAmount),
		unsafe.Offsetof(Smith.Name),
	)

	// Використання Alignof для визначення вирівнювання типів даних
	fmt.Printf("Вирівнювання Maried: %d, HasChildren: %d, Age: %d, BancAccountAmount: %d, Name: %d\n",
		unsafe.Alignof(Smith.Maried),
		unsafe.Alignof(Smith.HasChildren),
		unsafe.Alignof(Smith.Age),
		unsafe.Alignof(Smith.BancAccountAmount),
		unsafe.Alignof(Smith.Name),
	)

	// Прямий доступ до поля за допомогою unsafe
	int32Ptr := (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Smith)) + unsafe.Offsetof(Smith.Age)))
	*int32Ptr = 123
	fmt.Printf("Змінене значення Age: %d\n", Smith.Age)

	// Перетворення Pointer на *Person
	examplePtr := (*Person)(unsafe.Pointer(&Smith))
	examplePtr.BancAccountAmount = 1234567890
	fmt.Printf("Змінене значення BancAccountAmount: %d\n", Smith.BancAccountAmount)

	bytes := []byte("Hello, world!")
	str := BytesToString(bytes)
	fmt.Println(str)
}

// BytesToString конвертує слайс байтів у рядок без додаткового копіювання пам'яті.
func BytesToString(b []byte) string {
	// Отримання SliceHeader, який описує слайс
	var bh = (*reflect.SliceHeader)(unsafe.Pointer(&b))

	// Створення StringHeader, який описує рядок
	var sh = reflect.StringHeader{
		Data: bh.Data, // використовуємо ту ж адресу даних
		Len:  bh.Len,  // довжина слайса визначає довжину рядка
	}

	// Конвертація StringHeader назад у рядок без копіювання даних
	return *(*string)(unsafe.Pointer(&sh))
}
