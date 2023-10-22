package learn_golang_goroutine

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	var counter int64

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Jumlah : ", counter)
}
