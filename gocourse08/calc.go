package main

import (
	"fmt"
	"reflect"
)

func SumOfSquares(a, b any) (float64, error) {
	af, err := convert(a)
	if err != nil {
		return 0, err
	}

	bf, err := convert(b)
	if err != nil {
		return 0, err
	}

	return af*af + bf*bf, nil
}

func convert(num any) (float64, error) {
	switch v := num.(type) {
	case int:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	// ...
	default:
		return 0, fmt.Errorf("unsupported type: %s", reflect.TypeOf(num))
	}
}

type Number interface {
	~int
}

func SumOfSquaresG[T Number](a, b T) T {
	return a*a + b*b
}
