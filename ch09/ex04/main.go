package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var num int
	var err error
	if len(os.Args) < 2 {
		fmt.Println("illegal arg")
	} else {
		num, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("illegal arg:", err)
		}
	}

	chans := []chan struct{}{}
	for i := 0; i < num; i++ {
		chans = append(chans, make(chan struct{}))
	}

	var start time.Time
	go func() {
		start = time.Now()
		fmt.Println("Start:", start)
		chans[0] <- struct{}{}
		fmt.Println("send 0")
	}()

	for i := 0; i < num-1; i++ {
		i := i
		go func() {
			s := <-chans[i]
			fmt.Println("receive", i)
			chans[i+1] <- s
			fmt.Println("send", i+1)
		}()
	}

	<-chans[num-1]
	fmt.Println("receive", num-1)
	t := time.Now()
	fmt.Println("End:", t)
	fmt.Println(time.Since(start))
}
