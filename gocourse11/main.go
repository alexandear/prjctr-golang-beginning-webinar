package main

import (
	"fmt"

	"github.com/alexandear/prjctr-golang-beginning-webinar/gocourse11/pizza"
	"github.com/alexandear/prjctr-golang-beginning-webinar/gocourse11/service"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	builder()
	options()
	adapter()
	di()
}

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

type PaymentProcessor interface {
	Process(money int) error
}

type PaypalProcessor struct{}

func (p PaypalProcessor) Process(money int) error {
	fmt.Println("Paypal processing money", money)
	return nil
}

type StripeProcessor struct{}

func (p StripeProcessor) Withdraw(currency string, money int) error {
	fmt.Println("Stripe processing money", money, "in currency", currency)
	return nil
}

func process(p PaymentProcessor) {
	err := p.Process(100)
	if err != nil {
		fmt.Println(err)
	}
}

type StripeProcessorAdapter struct {
	StripeProcessor
}

func (s StripeProcessorAdapter) Process(money int) error {
	return s.Withdraw("USD", money)
}

func adapter() {
	paypal := PaypalProcessor{}
	stripe := StripeProcessorAdapter{StripeProcessor{}}

	process(paypal)
	process(stripe)
}

type PostgreSQL struct{}

func (PostgreSQL) User() (string, error) { return "", nil }

type MySQL struct{}

func (MySQL) User() (string, error) { return "", nil }

type ZapLogger struct{}

type Kafka struct{}

func di() {
	repository := &MySQL{}
	logger := &ZapLogger{}
	broker := &Kafka{}
	s := service.NewService(repository, logger, broker)
	_ = s.User()
}
