package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type linkedList struct {
	head *Node
}

func NewLinkedList() *linkedList {
	return &linkedList{head: nil}
}

func (l *linkedList) PushFront(value int) {
	newNode := &Node{value: value}
	newNode.next = l.head
	l.head = newNode
}

func (l *linkedList) Print() {
	current := l.head
	fmt.Println("head")
	for current != nil {
		fmt.Printf(" -> %d", current.value)
		current = current.next
	}

	fmt.Println(" -> nil")
}

func main() {
	list := NewLinkedList()
	list.PushFront(99)
	list.PushFront(30)
	list.PushFront(55)
	list.Print()

}
