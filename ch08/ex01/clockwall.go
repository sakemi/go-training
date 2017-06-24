package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	for _, arg := range os.Args[1:] {
		c := strings.Split(arg, "=")
		if len(c) != 2 {
			log.Printf("Illegal argument: %s", arg)
			continue
		}
		tz := c[0]
		url := c[1]
		fmt.Printf("%s\n", tz)
		go func() {
			conn, err := net.Dial("tcp", url)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			mustCopy(os.Stdout, conn)
		}()
	}
	for {
		fmt.Println("")
		time.Sleep(1 * time.Second)
	}
}

func writeClock(b []*bufio.Writer) {
	for _, buf := range b {
		buf.Flush()
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
