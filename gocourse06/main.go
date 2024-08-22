package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
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

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		task1()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		task2()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		task3()
	}()

	wg.Wait()

	elapsed := time.Since(now)
	fmt.Println("Elapsed:", elapsed)
}
