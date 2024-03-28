package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var _ BankClient = &Account{} // это трюк для проверки типа: до тех пор пока Account не будет реализовывать интерфейс BankClient, программа не запустится

type BankClient interface {
	// Deposit deposits given amount to clients account
	Deposit(amount int)

	// Withdrawal withdraws given amount from clients account.
	// return error if clients balance less the withdrawal amount
	Withdrawal(amount int) error

	// Balance returns clients balance
	Balance() int
}

// Структура для имплементации заданного интерфейса
type Account struct {
	bMutex  sync.RWMutex
	balance int
}

// Balance ...
func (a *Account) Balance() int {
	a.bMutex.RLock()
	defer a.bMutex.RUnlock()
	return a.balance
}

// Deposit ...
func (a *Account) Deposit(amount int) {
	a.bMutex.Lock()
	a.balance += amount
	a.bMutex.Unlock()
}

// Withdrawal ..
func (a *Account) Withdrawal(amount int) error {
	if amount > a.Balance() {
		return errors.New("your balance less then amount")
	}
	a.bMutex.Lock()
	a.balance -= amount
	a.bMutex.Unlock()
	return nil
}

// NewAccount ..
func NewAccount() *Account {
	return &Account{
		bMutex:  sync.RWMutex{},
		balance: 0,
	}
}
func main() {
	account := NewAccount()
	wg := sync.WaitGroup{}
	fmt.Println("Wait for preliminary operations")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Duration(rand.Intn(501)+500) * time.Millisecond)
			amount := rand.Intn(10) + 1
			account.Deposit(amount)
			wg.Done()
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Duration(rand.Intn(501)+500) * time.Millisecond)
			amount := rand.Intn(5) + 1
			if err := account.Withdrawal(amount); err != nil {
				fmt.Printf("Error while withdrawal operation: %s\n", err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Preliminary operation is ended.")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		commandline, _ := reader.ReadString('\n')
		commandline = strings.Replace(commandline, "\r\n", "", -1)
		commandArray := strings.Split(commandline, " ")
		switch commandArray[0] {
		case "balance":
			fmt.Printf("Your balance is %d\n", account.Balance())
		case "deposit":
			if len(commandArray) == 1 {
				fmt.Println("You have to enter amount sum after deposit command: deposit <amaunt>")
			} else {
				amaunt, err := strconv.Atoi(commandArray[1])
				if err != nil || amaunt < 0 {
					fmt.Println("Amaunt must be positive integer number")
				} else {
					account.Deposit(amaunt)
				}
			}
		case "withdrawal":
			if len(commandArray) == 1 {
				fmt.Println("You have to enter amount after withdrawal command: withdrawal <sum>")
			} else {
				amaunt, err := strconv.Atoi(commandArray[1])
				if err != nil || amaunt < 0 {
					fmt.Println("Amaunt must be positive integer number")
				} else {
					if err = account.Withdrawal(amaunt); err != nil {
						fmt.Printf("Error while withdrawal operation: %s\n", err)
					}
				}
			}
		case "exit":
			fmt.Println("Good bye.")
			return
		default:
			fmt.Println("Unsupported command. You can use commands: balance, deposit <amaunt>, withdrawal <amaunt>, exit")
		}
	}
}
