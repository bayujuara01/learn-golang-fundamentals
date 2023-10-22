package learn_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctxBg := context.Background()
	fmt.Println(ctxBg)

	ctxTd := context.TODO()
	fmt.Println(ctxTd)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "Key_B", "Data B")
	contextC := context.WithValue(contextA, "Key_C", "Data C")

	contextD := context.WithValue(contextB, "key_D", "Data D")
	contextE := context.WithValue(contextB, "Key_E", "Data E")

	contextF := context.WithValue(contextC, "Key_F", "Data F")
	contextG := context.WithValue(contextF, "Key_G", "Data G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println("==VALUE==")
	fmt.Println(contextF.Value("Key_B"))
	fmt.Println(contextF.Value("Key_C"))
	fmt.Println(contextE.Value("Key_E"))
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // simulate slow process
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())

	destination := CreateCounter(ctx)

	for i := range destination {
		fmt.Printf("Data : %d\n", i)
		if i == 10 {
			break
		}
	}

	cancel()

	time.Sleep(1 * time.Second)
	fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())

	destination := CreateCounter(ctx)
	fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())

	for i := range destination {
		fmt.Printf("Data : %d\n", i)
		if i == 10 {
			break
		}
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	parent := context.Background()
	futureTime := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(parent, futureTime)
	defer cancel()
	fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())

	destination := CreateCounter(ctx)
	fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())

	for i := range destination {
		fmt.Printf("Data : %d\n", i)
		if i == 10 {
			break
		}
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())
}
