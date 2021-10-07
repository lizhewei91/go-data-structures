package stack

import "fmt"

type Stack struct {
	data []int
}

// New: 返回stack的接口
func New() *Stack {
	return &Stack{
		data: []int{},
	}
}

// IsEmpty: 判断stack中是否有元素
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Peek: 返回栈顶的元素
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

// Push: 增加一个元素到stack
func (s *Stack) Push(n int) *Stack {
	s.data = append(s.data, n)
	return s
}

// Pop: 移除栈顶元素，并返回其值
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Stack is empty")
	}
	element := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return element, nil
}
