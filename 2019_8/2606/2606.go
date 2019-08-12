package main

import (
	"fmt"
)

type Node struct {
	edges   []*Node
	visited bool
}

type Stack struct {
	nodes []*Node
	count int
}

func main() {
	var numNodes, numEdges int
	stack := NewStack()
	fmt.Scanln(&numNodes)
	var node = make([]*Node, numNodes+1)
	for i := 0; i <= numNodes; i++ {
		node[i] = new(Node)
	}
	fmt.Scanln(&numEdges)

	for i := 0; i < numEdges; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		node[a].edges = append(node[a].edges, node[b])
		node[b].edges = append(node[b].edges, node[a])
	}

	stack.Push(node[1])

	for stack.count > 0 {
		temp := stack.Pop()
		stack.Visit(temp)
	}

	var count int = 0
	for i := 1; i <= numNodes; i++ {
		if node[i].visited {
			count++
		}
	}

	fmt.Println(count - 1)
}

func (s *Stack) Visit(n *Node) {
	for i := 0; i < len(n.edges); i++ {
		if !n.edges[i].visited {
			s.Push(n.edges[i])
		}
	}
}

func NewStack() *Stack {
	return &Stack{make([]*Node, 20), 0}
}

func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.Top().visited = true
	s.count--
	return s.nodes[s.count]
}

func (s *Stack) Top() *Node {
	return s.nodes[s.count-1]
}
