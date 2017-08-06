@echo off
go build
ex01.exe -t jpeg < in.png > png.jpg
ex01.exe -t gif < in.png > png.gif
ex01.exe -t png < in.jpg > jpeg.png
ex01.exe -t gif < in.jpg > jpeg.gif
ex01.exe -t jpeg < in.gif > gif.jpg
ex01.exe -t png < in.gif > gif.png
