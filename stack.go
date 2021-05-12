package gout

// StackNode struct
type StackNode struct {
	Val  interface{}
	Next *StackNode
}

// Stack struct
type Stack struct {
	Head *StackNode
}

// Peek method
func (s *Stack) Peek() *StackNode {
	return s.Head
}

// Pop method
func (s *Stack) Pop() *StackNode {
	if s.Head != nil {
		temp := s.Head
		s.Head = s.Head.Next
		temp.Next = nil
		return temp
	}

	return nil
}

// Push method
func (s *Stack) Push(n *StackNode) {
	n.Next = s.Head
	s.Head = n
}

// Len method
func (s *Stack) Len() int {
	node := s.Head
	len := 0

	for node != nil {
		len++
		node = node.Next
	}

	return len
}
