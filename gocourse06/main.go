package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func task1(done chan<- struct{}) {
	s := time.Duration((rand.N(300))+100) * time.Millisecond
	fmt.Println("Task 1 sleeping for", s)
	time.Sleep(s)
	fmt.Println("Task 1")
	done <- struct{}{}
}

func task2(ch chan<- struct{}) {
	s := time.Duration((rand.N(300))+100) * time.Millisecond
	fmt.Println("Task 2 sleeping for", s)
	time.Sleep(s)
	fmt.Println("Task 2")
	ch <- struct{}{}
}

func task3(done chan<- struct{}) {
	s := time.Duration((rand.N(300))+100) * time.Millisecond
	fmt.Println("Task 3 sleeping for", s)
	time.Sleep(s)
	fmt.Println("Task 3")
	done <- struct{}{}
}

func main() {
	now := time.Now()

	done := make(chan struct{})

	go task1(done)
	go task2(done)
	go task3(done)

	<-done
	<-done
	<-done

	elapsed := time.Since(now)
	fmt.Println("Elapsed:", elapsed)
}
