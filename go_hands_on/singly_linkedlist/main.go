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
	//fmt.Println(l.head)
}

func (l *LinkedList) PopFront() {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}
	l.head = l.head.next
}

func (l *LinkedList) Print() {
	temp := l.head
	for temp != nil {
		fmt.Printf("%d ->", temp.value)
		temp = temp.next
	}
	fmt.Println("nil")
}

func (l *LinkedList) PushBack(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode

}
func main() {
	list := NewLinkedList()
	list.PushFront(10)
	list.PushFront(20)
	list.PushBack(99)

	list.Print()
	list.PopFront()
	list.Print()

}
