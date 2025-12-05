package ds

import "errors"

type Stack struct {
	s []int
}

func (stack *Stack) Push(val int) {
	stack.s = append(stack.s, val)
}

func (stack *Stack) Pop() (int, error) {
	l := len(stack.s)
	if l == 0 {
		return 0, errors.New("Stack is empty")
	}
	res := stack.s[l-1]
	stack.s = stack.s[:l-1]
	return res, nil
}

func (stack *Stack) Len() int {
	return len(stack.s)
}

func (stack *Stack) Tail() int {
	return stack.s[stack.Len() - 1]
}

func (stack *Stack) S() []int {
	return stack.s
}

func NewStack() *Stack {
	return &Stack{make([]int, 0)}
}
