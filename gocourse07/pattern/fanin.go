//go:build fanin

package main

import (
	"fmt"
	"time"

	"math/rand/v2"
)

// the boring function return a channel to communicate with it.
func boring(msg string) <-chan string { // <-chan string means receives-only channel of string.
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.N(1000)) * time.Millisecond)
		}

	}()
	return c // return a channel to caller.
}

func fanInSimple(cs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, ci := range cs {
		go func() {
			for {
				v := <-ci
				c <- v
			}
		}() // send each channel to
	}
	return c
}

func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			v1 := <-c1
			c <- v1
		}
	}()
	go func() {
		for {
			v2 := <-c2 // read value from c2 and send it to c
			c <- v2
		}
	}()
	return c
}

func main() {
	// merge 2 channels into 1 channel
	// c := fanIn(boring("Joe"), boring("Ahn"))
	c := fanInSimple(boring("Joe"), boring("Ahn"))

	for v := range c {
		fmt.Println(v) // now we can read from 1 channel
	}
	fmt.Println("You're both boring. I'm leaving")
}
