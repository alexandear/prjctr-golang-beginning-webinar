package extend

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type DelayedCancelContext struct {
	parent     context.Context // Зберігаємо посилання на батьківський контекст
	cancelFunc context.CancelFunc
	mutex      sync.Mutex
	timer      *time.Timer
	delay      time.Duration
	doneChan   chan struct{}
}

func NewDelayedCancelContext(parent context.Context, delay time.Duration) *DelayedCancelContext {
	doneChan := make(chan struct{})
	cancelFunc := func() {
		close(doneChan)
	}

	dctx := &DelayedCancelContext{
		parent:     parent,
		cancelFunc: cancelFunc,
		delay:      delay,
		doneChan:   doneChan,
	}

	// Слухаємо сигнали від parent context для активації ScheduleDone
	go func() {
		<-parent.Done()
		dctx.ScheduleDone()
	}()

	return dctx
}

func (d *DelayedCancelContext) Deadline() (deadline time.Time, ok bool) {
	return d.parent.Deadline() // Передаємо виклик до батьківського контексту
}

func (d *DelayedCancelContext) Done() <-chan struct{} {
	return d.doneChan // Повертаємо наш власний канал
}

func (d *DelayedCancelContext) Err() error {
	select {
	case <-d.doneChan:
		return context.Canceled // Контекст скасовано
	default:
		return d.parent.Err() // Перевіряємо стан батьківського контексту
	}
}

func (d *DelayedCancelContext) Value(key any) any {
	return d.parent.Value(key) // Передаємо виклик до батьківського контексту
}

func (d *DelayedCancelContext) ScheduleDone() {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	go func() {
		i := d.delay / time.Second
		for {
			fmt.Printf("Stop after %d seconds\r", i)
			time.Sleep(1 * time.Second)
			i--
		}
	}()

	if d.timer != nil {
		d.timer.Stop()
	}
	d.timer = time.AfterFunc(d.delay, func() {
		d.cancelFunc()
	})
}
