package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		input := sc.Text()
		expr, err := Parse(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		env := Env{}
		for k, _ := range vars {
			var in string
			fmt.Printf("%s: ", k)
			fmt.Scan(&in)
			f, err := strconv.ParseFloat(in, 64)
			for err != nil {
				fmt.Println(err)
				fmt.Printf("%s: ", k)
				fmt.Scan(&in)
				f, err = strconv.ParseFloat(in, 64)
			}
			env[Var(k)] = f
		}
		fmt.Println(expr.Eval(env))
	}
}
