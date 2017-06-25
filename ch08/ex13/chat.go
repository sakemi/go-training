package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const timeout = 1 //min

type client struct {
	ch   chan<- string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			for c := range clients {
				cli.ch <- c.name + " is in this room."
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	cli := client{ch, who}
	entering <- cli

	write := make(chan struct{})
	go disconnecter(conn, write)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		write <- struct{}{}
		messages <- who + ": " + input.Text()
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func disconnecter(conn net.Conn, ch <-chan struct{}) {
	cnt := 0
	tick := time.Tick(1 * time.Minute)
	for cnt < timeout {
		select {
		case <-ch:
			cnt = 0
		case <-tick:
			cnt++
		}
	}
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
