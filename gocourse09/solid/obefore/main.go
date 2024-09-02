// The program violates the Open/Closed Principle from SOLID.

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type calculator struct{}

func (a calculator) areaSum(shapes ...any) float64 {
	var sum float64
	for _, shape := range shapes {
		switch sh := shape.(type) {
		case circle:
			r := sh.radius
			sum += math.Pi * r * r
		case square:
			l := sh.length
			sum += l * l
		}
	}
	return sum
}

func main() {
	c := circle{radius: 5}
	s := square{length: 7}
	var calc calculator
	fmt.Println("area sum:", calc.areaSum(c, s))
}
