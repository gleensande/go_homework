package main

import (
	"os"
	"io"
	"bufio"
	"strconv"
	"strings"
	"errors"
)

func main() {
	out := os.Stdout
	in := os.Stdin

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		calc(scanner.Text(), out)
	}
}

func calc(in string, out io.Writer) error {
	ins := strings.Split(in, " ")
	var stack [1000]int
	var err error

	sp := 0

	for _, r := range ins {
		var x int

		if sp < 0 {
			return errors.New("abnormal syntax")
		}

		switch r {
			case " " : fallthrough
			case "\n" :
			case "=" :
				break
			case "+":
				if sp < 2 {
					return errors.New("abnormal syntax")
				}
				stack[sp - 2] = stack[sp - 2] + stack[sp - 1]
				sp--
			case "-":
				if sp < 2 {
					return errors.New("abnormal syntax")
				}
				stack[sp - 2] = stack[sp - 2] - stack[sp - 1]
				sp--
			case "*":
				if sp < 2 {
					return errors.New("abnormal syntax")
				}
				stack[sp - 2] = stack[sp - 1] * stack[sp - 2]
				sp--
			case "/":
				if sp < 2 {
					return errors.New("abnormal syntax")
				}
				stack[sp - 2] = stack[sp - 2] / stack[sp - 1]
				sp--
			default:
				x, err = strconv.Atoi(r)
				if err != nil {
					return err
				}
				stack[sp] = x
				sp++
		}
	}

	out.Write([]byte(strconv.Itoa(stack[sp - 1])))

	return err
}
