// The program follows the Single Responsibility Principle from SOLID.

package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) name() string {
	return "circle"
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	length float64
}

func (s square) name() string {
	return "square"
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
	name() string
}

type outputter struct{}

func (o outputter) Text(s shape) string {
	return fmt.Sprintf("area of the %s: %f", s.name(), s.area())
}

func (o outputter) JSON(s shape) (string, error) {
	res := struct {
		Name string  `json:"shape"`
		Area float64 `json:"area"`
	}{
		Name: s.name(),
		Area: s.area(),
	}

	bs, err := json.Marshal(res)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func main() {
	c := circle{radius: 5}
	s := square{length: 7}
	var out outputter

	fmt.Println("Text output:")
	fmt.Println(out.Text(c))
	fmt.Println(out.Text(s))

	fmt.Println("\nJSON output:")
	fmt.Println(out.JSON(s))
	fmt.Println(out.JSON(c))
}
