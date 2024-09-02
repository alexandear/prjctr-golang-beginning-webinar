package main

import "fmt"

func CompareAny[T comparable](a, b T) bool {
	return a == b
}

type IntOrFloat interface {
	int | float64
}

func SumNumbers[T IntOrFloat](numbers []T) T {
	var sum T
	for _, num := range numbers {
		sum += num
	}
	return sum
}

type GenericList[T any] struct {
	elements []T
}

func (g *GenericList[T]) Add(element T) {
	g.elements = append(g.elements, element)
}

func (g *GenericList[T]) All() []T {
	return g.elements
}

func min[T int | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func PtrOf(v int) *int {
	return &v
}

func main() {
	fmt.Println(PtrOf())
}
