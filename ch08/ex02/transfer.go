package main

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func retrData(ctrl *Ctrl, addr string, file string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
		ctrl.reply(failToConn)
		return
	}
	defer conn.Close()

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		ctrl.reply(failToOpenFile)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		log.Println(err)
		ctrl.reply(transferAbort)
		return
	}

	ctrl.reply(closeConn)
}

func storData(ctrl *Ctrl, addr string, file string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
		ctrl.reply(failToConn)
		return
	}
	defer conn.Close()

	f, err := os.Create(file)
	if err != nil {
		log.Println(err)
		ctrl.reply(failToOpenFile)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, conn)
	if err != nil {
		log.Println(err)
		ctrl.reply(transferAbort)
		return
	}

	ctrl.reply(transferComplete)
}
