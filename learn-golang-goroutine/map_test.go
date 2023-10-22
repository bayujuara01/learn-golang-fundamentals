package learn_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

type OwnMap struct {
	syncMap   *sync.Map
	waitGroup *sync.WaitGroup
}

type Data struct {
	key string
	id  interface{}
}

func (ownMap *OwnMap) AddToMap(data *Data) {
	defer ownMap.waitGroup.Done()
	ownMap.syncMap.Store(data.key, data.id)
}

func TestMap(t *testing.T) {
	waitGroup := new(sync.WaitGroup)
	syncMap := new(sync.Map)
	ownMap := OwnMap{syncMap, waitGroup}

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		key := fmt.Sprintf("Key_%d", i)
		go ownMap.AddToMap(&Data{key, i})
	}

	ownMap.waitGroup.Wait()
	ownMap.syncMap.Range(func(key, value any) bool {
		fmt.Printf("Key : %s, ID : %d, Type : %T\n", key, value, value)
		return true
	})
	fmt.Println("Selesai")
}
