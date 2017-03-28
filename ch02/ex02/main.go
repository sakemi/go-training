package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			args = append(args, input.Text())
		}
	}
	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))

		feet := Feet(t)
		meters := Meters(t)
		fmt.Printf("%s = %s, %s = %s\n",
			feet, FToM(feet), meters, MToF(meters))

		kg := Kg(t)
		lb := Lb(t)
		kan := Kan(t)
		fmt.Printf("%s = %s = %s, %s = %s = %s, %s = %s = %s\n",
			kg, KgToLb(kg), KgToKan(kg), lb, LbToKg(lb), LbToKan(lb), kan, KanToKg(kan), KanToLb(kan))
	}
}
