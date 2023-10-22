package learn_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Sync struct {
	pool      *sync.Pool
	waitGroup *sync.WaitGroup
}

func (sn *Sync) GetFromPool() {
	defer sn.waitGroup.Done()
	data := sn.pool.Get()
	fmt.Println(data)
	time.Sleep(1 * time.Second)
	sn.pool.Put(data)
}

func TestPool(t *testing.T) {
	sn := &Sync{
		pool:      &sync.Pool{},
		waitGroup: &sync.WaitGroup{},
	}

	sn.pool.Put("Hello")
	sn.pool.Put("Bayu Seno Ariefyanto")
	sn.pool.Put("Belajar Terus")

	for i := 0; i < 10; i++ {
		sn.waitGroup.Add(1)
		go sn.GetFromPool()
	}

	sn.waitGroup.Wait()
	fmt.Println("Selesai")

}

func TestPoolDefault(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} { return "New Value" },
	}
	sn := &Sync{
		pool:      pool,
		waitGroup: new(sync.WaitGroup),
	}

	sn.pool.Put("Hello")
	sn.pool.Put("Bayu Seno Ariefyanto")
	sn.pool.Put("Belajar Terus")

	for i := 0; i < 10; i++ {
		sn.waitGroup.Add(1)
		go sn.GetFromPool()
	}

	sn.waitGroup.Wait()
	fmt.Println("Selesai")
}
