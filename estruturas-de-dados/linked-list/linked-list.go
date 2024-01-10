package linkedlist

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Tail *Node
	Head *Node
}

func (ll *LinkedList) Append(data interface{}) *LinkedList {
	nextNode := &Node{
		Next: nil,
		Data: data,
	}

	if ll.Head.Data == nil {
		ll.Head = nextNode
	}

	if ll.Head.Data != nil {
		ll.Tail.Next = nextNode
	}

	ll.Tail = nextNode

	return ll
}

func (ll *LinkedList) DeleteWithValue(value interface{}) *LinkedList {
	currentNode := ll.Head

	if currentNode.Data == value {
		ll.Head = ll.Head.Next
		return ll
	}

	for {
		if value == currentNode.Next.Data {
			if currentNode.Next.Next != nil {
				currentNode.Next = currentNode.Next.Next
				return ll
			}

			currentNode.Next = nil
			return ll
		}

		currentNode = currentNode.Next
	}
}

func (ll *LinkedList) PrintAll() {
	currentNode := ll.Head

	for {
		fmt.Println(currentNode.Data)
		if currentNode.Next == nil {
			return
		}
		currentNode = currentNode.Next
	}
}

func New() *LinkedList {
	emptyNode := &Node{
		Data: nil,
		Next: nil,
	}

	return &LinkedList{
		Tail: emptyNode,
		Head: emptyNode,
	}
}
