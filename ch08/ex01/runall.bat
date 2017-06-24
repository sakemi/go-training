@echo off
go build clock.go
go build clockwall.go
start clock.exe -port 8000
start clock.exe -port 8010
start clock.exe -port 8020
clockwall.exe JP=localhost:8000 UTC=localhost:8010 US=localhost:8020
