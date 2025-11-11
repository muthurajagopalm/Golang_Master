package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

func (l *LinkedList) PushFront(value int) {
	newNode := &Node{value: value}
	newNode.next = l.head
	l.head = newNode

}

func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ->", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func (l *LinkedList) PushBack(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *LinkedList) PopFront() {
	if l.head == nil {
		fmt.Println("The list is empty and nothing to remove")
		return

	}

	l.head = l.head.next
}

func (l *LinkedList) Length() int {
	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func main() {
	list := NewLinkedList()
	list.PushFront(5)
	list.PushBack(10)
	list.PopFront()
	list.Length()
	list.Print()

}
