package learn_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	tick := time.NewTicker(500 * time.Millisecond)

	go func() {
		time.Sleep(5 * time.Second)
		tick.Stop()
	}()

	for tk := range tick.C {
		fmt.Println(tk)
	}
}

func TestTick(t *testing.T) {
	tick := time.Tick(2 * time.Second)

	for tk := range tick {
		fmt.Println(tk)
	}
}
