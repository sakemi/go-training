package main

var deposits = make(chan int)
var balances = make(chan int)
var withdraw = make(chan int)
var withdrawResult = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	withdraw <- amount
	return <-withdrawResult
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraw:
			if balance < amount {
				withdrawResult <- false
			} else {
				balance -= amount
				withdrawResult <- true
			}
		}
	}
}

func init() {
	go teller()
}
