package learn_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
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

type IBankAccount interface {
	SetBalance(newBalance int)
	AddBalance(amount int)
	GetBalance() int
}

type BankAccount struct {
	RWMutex sync.RWMutex
	balance int
}

func (bankAccount *BankAccount) SetBalance(newBalance int) {
	bankAccount.RWMutex.Lock()
	bankAccount.balance = newBalance
	bankAccount.RWMutex.Unlock()
}

func (bankAccount *BankAccount) AddBalance(amount int) {
	bankAccount.RWMutex.Lock()
	bankAccount.balance = bankAccount.balance + amount
	bankAccount.RWMutex.Unlock()
}

func (bankAccount *BankAccount) GetBalance() int {
	bankAccount.RWMutex.RLock()
	balance := bankAccount.balance
	bankAccount.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	bankAccount := new(BankAccount)

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				bankAccount.AddBalance(1)
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance : ", bankAccount.GetBalance())
}
