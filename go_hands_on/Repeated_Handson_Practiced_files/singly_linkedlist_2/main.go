package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) PushFront(value int) {
	newNode := &Node{value: value}
	newNode.next = l.head
	l.head = newNode
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

	fmt.Printf("Removing %d from the front\n", l.head.value)
	l.head = l.head.next
}

func (l *LinkedList) PopBack() {

	if l.head == nil {
		fmt.Println("The list is empty and nothing to remove")
		return
	}

	if l.head.next != nil {
		fmt.Printf("Removing %d value from the back\n", l.head.value)
		l.head = nil
		return
	}

	current := l.head
	for current.next.next != nil {
		current = current.next
	}
	fmt.Printf("Removing %d from the back\n", current.next.value)
	current.next = nil

}

func (l *LinkedList) Print() {
	current := l.head

	if current == nil {
		fmt.Println("The list is empty and nothing to remove")
	}

	for current != nil {
		fmt.Printf("%d ->", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func (l *LinkedList) Length() int {
	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.next
	}
	fmt.Println("The length of the node count", count)
	return count
}

func (l *LinkedList) Find(value int) bool {
	current := l.head

	for current != nil {
		if current.value == value {
			return true

		}
		current = current.next
	}
	return false

}

func main() {
	list := &LinkedList{}
	list.PushFront(10)
	list.PushFront(100)
	list.PushFront(200)
	list.PushFront(500)
	list.PushFront(2000)
	list.PushBack(10000)
	list.Length()
	list.Find(500)
	list.Print()
}
