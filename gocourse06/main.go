package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func task1() {
	s := time.Duration((rand.N(300))+100) * time.Millisecond
	fmt.Println("Task 1 sleeping for", s)
	time.Sleep(s)
	fmt.Println("Task 1")
}

func task2() {
	s := time.Duration((rand.N(300))+100) * time.Millisecond
	fmt.Println("Task 2 sleeping for", s)
	time.Sleep(s)
	fmt.Println("Task 2")
}

func task3() {
	s := time.Duration((rand.N(300))+100) * time.Millisecond
	fmt.Println("Task 3 sleeping for", s)
	time.Sleep(s)
	fmt.Println("Task 3")
}

func main() {
	now := time.Now()

	task1()
	task2()
	task3()

	elapsed := time.Since(now)
	fmt.Println("Elapsed:", elapsed)
}
