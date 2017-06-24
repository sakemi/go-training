@echo off
go build reverb.go
go build netcat.go
start reverb.exe
netcat.exe < input.txt
