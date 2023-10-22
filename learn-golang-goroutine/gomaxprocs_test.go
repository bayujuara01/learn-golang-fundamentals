package learn_golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	waitGroup := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			time.Sleep(3 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	waitGroup.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	waitGroup := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			time.Sleep(3 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	runtime.GOMAXPROCS(1)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	waitGroup.Wait()
}
