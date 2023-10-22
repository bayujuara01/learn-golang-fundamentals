package learn_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type SyncCond struct {
	cond      *sync.Cond
	waitGroup *sync.WaitGroup
}

func (syncCond *SyncCond) WaitCondition(value int) {
	defer func() {
		syncCond.waitGroup.Done()
		syncCond.cond.L.Unlock()
	}()

	syncCond.cond.L.Lock()
	syncCond.cond.Wait()
	fmt.Println("Done : ", value)
}

func TestCond(t *testing.T) {
	locker := new(sync.Mutex)
	waitGroup := new(sync.WaitGroup)
	cond := sync.NewCond(locker)
	sc := SyncCond{
		cond:      cond,
		waitGroup: waitGroup,
	}

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go sc.WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			cond.Signal()
		}
	}()

	waitGroup.Wait()

	fmt.Println("Selesai")
}

func TestCondBroadcast(t *testing.T) {
	locker := new(sync.Mutex)
	waitGroup := new(sync.WaitGroup)
	cond := sync.NewCond(locker)
	sc := SyncCond{
		cond:      cond,
		waitGroup: waitGroup,
	}

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go sc.WaitCondition(i)
	}

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	waitGroup.Wait()

	fmt.Println("Selesai")
}
