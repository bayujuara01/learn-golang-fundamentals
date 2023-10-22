package learn_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunOnTime(timer <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-timer)
}

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	wg := sync.WaitGroup{}

	fmt.Println(time.Now())

	wg.Add(1)
	go RunOnTime(timer.C, &wg)
	wg.Wait()
}

func TestAfter(t *testing.T) {
	fmt.Println(time.Now())
	timeAfter := time.After(3 * time.Second)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go RunOnTime(timeAfter, &wg)

	wg.Wait()
}

func TestAfterFunc(t *testing.T) {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	fmt.Println(time.Now())
	time.AfterFunc(3*time.Second, func() {
		defer waitGroup.Done()
		fmt.Println(time.Now())
	})

	waitGroup.Wait()
}
