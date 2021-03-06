package gout

import (
	"fmt"
)

// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedList struct
type LinkedList struct {
	head *ListNode
}

// Head method
func (ll *LinkedList) Head() *ListNode {
	return ll.head
}

// String method
func (ll *LinkedList) String() string {
	node := ll.head
	str := ""

	for node != nil {
		str += fmt.Sprintf("%s -> %d", str, node.Val)
		node = node.Next
	}

	str += "nil"

	return str
}

// String method
func (ll *ListNode) String() string {
	str := ""
	node := ll

	for node != nil {
		str += fmt.Sprintf("%s -> %d", str, node.Val)
		node = node.Next
	}

	str += "nil"

	return str
}

// GetLinkedListFromArray method
func GetLinkedListFromArray(array []int) *LinkedList {
	ll := &LinkedList{}

	if len(array) == 0 {
		return ll
	}

	root := &ListNode{Val: array[0], Next: nil}
	cur := root

	for i := 1; i < len(array); i++ {
		node := &ListNode{Val: array[i], Next: nil}
		cur.Next = node
		cur = node
	}

	ll.head = root

	return ll
}

// NewLinkedListFromArray method
func NewLinkedListFromArray(array []int) *LinkedList {
	ll := &LinkedList{}

	if len(array) == 0 {
		return ll
	}

	root := &ListNode{Val: array[0], Next: nil}
	cur := root

	for i := 1; i < len(array); i++ {
		node := &ListNode{Val: array[i], Next: nil}
		cur.Next = node
		cur = node
	}

	return ll
}
