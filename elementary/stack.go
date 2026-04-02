package elementary

import "fmt"

type Stack struct {
	Top   int
	Items []int
}

func RunStack() {
	stack := MakeStack(6)
	stack.Push(4)
	stack.Push(1)
	stack.Push(3)
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
	stack.Push(8)
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
	fmt.Println(&stack.Items)
}

func MakeStack(size int) *Stack {
	return &Stack{
		Top:   0,
		Items: make([]int, size),
	}
}

func (s *Stack) Empty() bool {
	return s.Top == 0
}

func (s *Stack) Push(i int) {
	s.Items[s.Top] = i
	s.Top += 1
}

func (s *Stack) Pop() int {
	s.Top -= 1
	return s.Items[s.Top+1]
}
