package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/stackskill0/LearningGo/Chapter4/stack"
)

func executeIt(st stack.Stack, op string) (ret int, err error) {

	p, _ := st.Pop()
	q, _ := st.Pop()
	switch op {
	case "+":
		ret = q + p
	case "-":
		ret = q - p
	case "*":
		ret = q * p
	case "/":
		ret = q / p
	default:
		ret = 0
		err = errors.New("Bad input")
	}
	return ret, err
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("RPN:")

	st := new(stack.Stack)

	for {
		s, err := reader.ReadString('\n')
		var token string
		if err != nil {
			return
		}
		for _, c := range s {
			switch {
			case c >= '0' && c <= '9':
				token = token + string(c)
			case c == ' ':
				r, _ := strconv.Atoi(token)
				st.Push(r)
				token = ""
			case c == '+' || c == '-' || c == '*' || c == '/':
				result, error := executeIt(*st, string(c))
				if error == nil {
					fmt.Printf("%d\n", result)
				} else {
					fmt.Println(error, result)
				}

			case c == 'q':
				return

			}
		}
	}
}
