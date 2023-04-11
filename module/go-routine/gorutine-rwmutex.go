package main

import (
	"fmt"
	"sync"
	"time"
	// "time"
)

type BackAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BackAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BackAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func main() {
	account := BackAccount{}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance :", account.GetBalance())
}
