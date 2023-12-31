package learn_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	counter := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				counter = counter + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Jumlah : ", counter)
}
