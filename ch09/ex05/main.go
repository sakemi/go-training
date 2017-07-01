package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	timeup := make(chan struct{})
	done := make(chan struct{})

	go func() {
		time.Sleep(1 * time.Second)
		close(timeup)
	}()

	go func() {
		ball := struct{}{}
		cnt := 0
	LOOP:
		for {
			select {
			case <-timeup:
				break LOOP
			default:
			}
			cnt++
			fmt.Println("ping")
			ch <- ball
			<-ch
		}
		fmt.Println(cnt) //30,000ぐらい
		close(done)
	}()

	go func() {
		ball := struct{}{}
	LOOP:
		for {
			select {
			case <-timeup:
				break LOOP
			default:
			}
			<-ch
			fmt.Println("pong")
			ch <- ball
		}
	}()

	<-done
}
