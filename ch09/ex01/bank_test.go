package main

import (
	"fmt"
	"testing"
)

func TestDeposit(t *testing.T) {
	done := make(chan struct{})

	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	// test withdraw ----------------------------------
	// withdraw OK
	wd := make(chan bool)
	go func() {
		res := Withdraw(200)
		fmt.Println("=", Balance())
		wd <- res
	}()
	go func() {
		res := Withdraw(100)
		wd <- res
	}()

	res := <-wd
	if !res {
		t.Errorf("Failed to withdraw.")
	}
	res = <-wd
	if !res {
		t.Errorf("Failed to withdraw.")
	}
	if got, want := Balance(), 0; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	// cannot withdraw because of balance shortage
	go func() {
		Deposit(50)
		done <- struct{}{}
	}()
	go func() {
		res := Withdraw(200)
		wd <- res
	}()
	go func() {
		res := Withdraw(100)
		wd <- res
	}()

	<-done
	res = <-wd
	if res {
		t.Errorf("Withdrew even if balance is not enouth")
	}
	res = <-wd
	if res {
		t.Errorf("Withdrew even if balance is not enouth")
	}
	if got, want := Balance(), 50; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
	// test withdraw end ----------------------------------
}
