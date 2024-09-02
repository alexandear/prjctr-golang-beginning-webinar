// The program follows the Liskov Substitutions Principle from SOLID.

package main

import "fmt"

type human struct {
	hname string
}

func (h human) name() string {
	return h.hname
}

type teacher struct {
	human
	degree string
	salary float64
}

type student struct {
	human
	grades map[string]int
}

type person interface {
	name() string
}

type printer struct{}

func (printer) info(p person) {
	fmt.Println("Name:", p.name())
}

func main() {
	h := human{hname: "Alex"}
	s := student{
		human: human{hname: "Mike"},
		grades: map[string]int{
			"Math":    8,
			"English": 9,
		},
	}
	t := teacher{
		human:  human{hname: "John"},
		degree: "CS",
		salary: 2000,
	}

	var p printer
	p.info(h)
	p.info(s)
	p.info(t)
}
