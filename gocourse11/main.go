package main

import (
	"github.com/alexandear/prjctr-golang-beginning-webinar/gocourse11/pizza"
	"github.com/davecgh/go-spew/spew"
)

type Director struct {
	builder pizza.Builder
}

func NewDirector(builder pizza.Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() *pizza.Pizza {
	return d.builder.
		SetDough("Thin Crust").
		SetSauce("Tomato").
		SetCheese("Mozzarella").
		Build()
}

func builder() {
	// dominosBuilder := pizza.NewDominosBuilder()
	ilMolinoBuilder := pizza.NewIlMolinoBuilder()

	// director := NewDirector(dominosBuilder)
	director := NewDirector(ilMolinoBuilder)
	pizza := director.Construct()

	spew.Dump(pizza)
}

func options() {
	pizza := pizza.NewPizza(
		pizza.WithDough("Thin Crust"),
		pizza.WithSauce("Tomato"),
		pizza.WithCheese("Mozzarella"),
	)

	spew.Dump(pizza)
}

func adapter() {
}

func main() {
	builder()
	options()
}
