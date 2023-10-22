package learn_golang_goroutine

import (
	"fmt"
	"testing"
)

func HelloGoroutine() {
	fmt.Println("Hello, Goroutine")
}

func TestHelloGoroutine(t *testing.T) {
	go HelloGoroutine()

	//time.Sleep(3 * time.Second)

	fmt.Println("Running")
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		DisplayNumber(i)
	}

	//time.Sleep(5 * time.Second)
}

//func BenchmarkManyGoroutine(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		go DisplayNumber(i)
//	}
//}
