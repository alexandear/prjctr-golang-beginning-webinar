// Треба змоделювати як сигнал від мозку надходить до кінцівок людини,
// якщо та танцює (достатньо ніг).
// У випадку, коли людина сідає, сигнал припиняє надходити.
// А як знову починає танцювати — сигнал знову починає надходити.
// Треба опрацювати закриття каналу і відключення воркерів,
// які читали з цього каналу. І знову запуск читання з каналу воркерами.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func brain(ctx context.Context, signal chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Man is gone")
			return
		case <-signal:
			fmt.Println("Dancing")
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	signal := make(chan struct{}, 2)

	// signal
	signal <- struct{}{}

	time.Sleep(2 * time.Second)

	signal <- struct{}{}

	ctx, cancel := context.WithTimeout(context.Background(), 64*60*time.Second)
	cancel()
	go brain(ctx, signal, &wg)

	wg.Wait()
}
