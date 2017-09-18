#!/bin/sh
chmod u+x runall.sh
go build bzip2.go
echo hoge > hoge.txt
./bzip2 < hoge.txt
