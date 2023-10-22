package learn_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	fmt.Println("Hello Wait Group")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	waitGroup := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(waitGroup)
	}

	waitGroup.Wait()
	fmt.Println("Selesai")
}
