package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// сюда писать код
// фукция main тоже будет тут

type Stack struct {
	Data []int
	Last int
	Size int
}

func NewStack() *Stack{
	return &Stack{make([]int, 0), 0, 0}
}

func (s *Stack) Push(element int) {
	s.Data = append(s.Data, element)
	s.Last = element
	s.Size++
}

func (s *Stack) Pop() int {
	if s.Size == 0 {
		panic("Stack is empty")
	}

	s.Size--
	element := s.Last
	if s.Size == 0 {
		s.Last = -1
	} else {
		s.Last = s.Data[s.Size - 1]
	}
	s.Data = s.Data[:s.Size]

	return element
}


func main()  {
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err == nil {
		fmt.Println(line)
	}


	//fmt.Printf("Result = %d\n", Calculate(line))
}


func Calculate(expression string) (int, error) {
	stack := NewStack()

	elements := strings.Split(expression, " ")


	for _, char := range elements {
		switch char {
		case " ":
			continue
		case "=":
			return stack.Pop(), nil
		case "+":
			stack.Push(stack.Pop() + stack.Pop())
		case "-":
			stack.Push(-stack.Pop() + stack.Pop())
		case "*":
			stack.Push(stack.Pop() * stack.Pop())
		case "/":
			stack.Push(stack.Pop() / stack.Pop())
		default:
			value, err := strconv.Atoi(char)
			if err != nil {
				return 0, err
			}

			fmt.Println(value)
			stack.Push(value)
		}
	}

	return 0, errors.New("not enough operations in input string")
}