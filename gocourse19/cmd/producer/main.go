package main

import (
	"context"
	"log"

	"gocourse19/pkg/produce"
)

func main() {
	pp := produce.NewPool(
		context.Background(),
		1,
		`main`,
		`amqp://localhost`,
		`guest`,
		`guest`,
	)

	for _, p := range pp.Producers() {
		if err := p.Push(
			context.Background(),
			`key-product`,
			[]byte(`{"someField1": "Some Value 1"}`),
		); err != nil {
			log.Println(err)
		}

		if err := p.Push(
			context.Background(),
			`key-brand`,
			[]byte(`{"someField2": "Some Value 2"}`),
		); err != nil {
			log.Println(err)
		}
	}
}
