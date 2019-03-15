package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	Stack    []int
	RealSize int
}

func (s *Stack) init() {
	s.RealSize = 0
}

func (s *Stack) push(n int) {
	if s.RealSize < len(s.Stack) {
		s.Stack[s.RealSize] = n
	} else {
		s.Stack = append(s.Stack, n)
	}
	s.RealSize++
}

func (s *Stack) pop() (int, error) {
	if s.RealSize <= 0 {
		return 0, errors.New("stack is empty")
	}
	s.RealSize--
	return s.Stack[s.RealSize], nil
}

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

	var err error
	abnormErr := errors.New("abnormal syntax")

	var stack Stack
	stack.init()

	for _, r := range ins {
		var x int
		var popped1, popped2 int

		switch r {
		case " ":
			fallthrough
		case "\n":
		case "=":
			break
		case "+":
			popped1, err = stack.pop()
			if err != nil {
				return abnormErr
			}

			popped2, err = stack.pop()
			if err != nil {
				return abnormErr
			}

			stack.push(popped2 + popped1)
		case "-":
			popped1, err = stack.pop()
			if err != nil {
				return abnormErr
			}

			popped2, err = stack.pop()
			if err != nil {
				return abnormErr
			}

			stack.push(popped2 - popped1)
		case "*":
			popped1, err = stack.pop()
			if err != nil {
				return abnormErr
			}

			popped2, err = stack.pop()
			if err != nil {
				return abnormErr
			}

			stack.push(popped2 * popped1)
		case "/":
			popped1, err = stack.pop()
			if err != nil {
				return abnormErr
			}
			if popped1 == 0 {
				return errors.New("division by 0")
			}

			popped2, err = stack.pop()
			if err != nil {
				return abnormErr
			}

			stack.push(popped2 / popped1)
		default:
			x, err = strconv.Atoi(r)
			if err != nil {
				return err
			}
			stack.push(x)
		}
	}

	popped, err := stack.pop()
	if err != nil {
		return abnormErr
	}
	out.Write([]byte(strconv.Itoa(popped)))

	return err
}
