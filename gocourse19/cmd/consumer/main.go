package main

import (
	"context"
	"encoding/json"
	"gocourse19/pkg/consume"
	"log"
	"time"
)

func main() {
	cp := consume.NewPool(
		context.Background(),
		1,
		2,
		`q4`,
		`amqp://localhost`,
		`guest`,
		`guest`,
	)

	for _, c := range cp.Consumers() {
		if err := c.InitStream(context.Background()); err != nil {
			log.Println(err)
		}

		for {
			for {
				if !c.IsDeliveryReady {
					log.Println(`Waiting...`)
					time.Sleep(cp.ChannelReconnectDelay())
				} else {
					break
				}
			}

			select {
			case d := <-c.GetStream():
				res := make(map[string]any)
				if err := json.Unmarshal(d.Body, &res); err != nil {
					log.Println(err)
				}
				log.Println(res)

				if err := d.Ack(false); err != nil {
					log.Println(err)
				}
			}
		}
	}
}
