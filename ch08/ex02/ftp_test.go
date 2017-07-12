package main

import (
	"fmt"
	"net"
	"os"
	"testing"
	"time"
)

type ConnMock struct{}

func (c ConnMock) Read(b []byte) (n int, err error) {
	return 0, fmt.Errorf("mock")
}
func (c ConnMock) Write(b []byte) (n int, err error) {
	return os.Stdout.Write(b)
}
func (c ConnMock) Close() error {
	return fmt.Errorf("mock")
}
func (c ConnMock) LocalAddr() net.Addr {
	return nil
}
func (c ConnMock) RemoteAddr() net.Addr {
	return nil
}
func (c ConnMock) SetDeadline(t time.Time) error {
	return fmt.Errorf("mock")
}
func (c ConnMock) SetReadDeadline(t time.Time) error {
	return fmt.Errorf("mock")
}
func (c ConnMock) SetWriteDeadline(t time.Time) error {
	return fmt.Errorf("mock")
}

func TestInterpret(t *testing.T) {
	conMcok := ConnMock{}
	ctrl := Ctrl{conMcok}
	handleRequests := func(reqs []string, info *Info) {
		for _, req := range reqs {
			if !interpret(&ctrl, info, req) {
				t.Error(req)
			}
		}
	}

	requests := []string{
		"USER username",
		"PASS password",
		"CWD /dir1/",
		"QUIT",
		"PORT 127,0,0,1,31,144",
		"TYPE A",
		"STRU struct",
		"RETR test.txt",
		"STOR test.txt",
		"NOOP",
	}
	info := &Info{}
	handleRequests(requests, info)

	//非サポートコマンド
	req := "HOGE foo bar"
	if interpret(&ctrl, info, req) {
		t.Error(req)
	}
}
