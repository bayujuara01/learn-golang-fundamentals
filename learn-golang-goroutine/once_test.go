package learn_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func RunThisOnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := &sync.Once{}
	waitGroup := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func() {
			//RunThisOnlyOnce()
			once.Do(RunThisOnlyOnce)
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
	fmt.Printf("Selesai, Counter : %d\n", counter)
}
