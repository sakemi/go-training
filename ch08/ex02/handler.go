package main

import (
	"bufio"
	"fmt"
	"log"
)

func handleCtrl(c Ctrl) {
	defer c.Close()
	log.Println("Connected.")
	c.reply(newUser)

	info := &Info{dir: "."}
	input := bufio.NewScanner(c)
	for input.Scan() {
		command := input.Text()
		fmt.Println(command)
		interpret(&c, info, command)
	}
}
