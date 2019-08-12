package main

import (
	"fmt"
)

type Node struct {
	value int64
}

type Stack struct {
	nodes []*Node
	count int64
}

func main() {
	stack := NewStack()
	var t int
	fmt.Scanln(&t)

	var sum int64
	for t > 0 {
		var num int64
		fmt.Scanln(&num)
		if num == 0 {
			stack.Pop()
		} else {
			stack.Push(&Node{num})
		}
		t--
	}

	for stack.count > 0 {
		sum += stack.Pop().value
	}

	fmt.Println(sum)
}

func NewStack() *Stack {
	return &Stack{make([]*Node, 1000000), 0}
}

func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}
