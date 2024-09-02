//go:build generator

package main

import (
	"fmt"
	"time"

	"math/rand/v2"
)

func main() {
	c := generator("boring!")
	for v := range c {
		fmt.Printf("You say: %q\n", v)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func generator(msg string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for i := 0; i < 3; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.N(1000)) * time.Millisecond)
		}
	}()
	return c
}
