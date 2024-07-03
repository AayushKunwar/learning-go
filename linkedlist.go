package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(data int) {
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
}

func (l *LinkedList) InsertTail(data int) {
	newNode := &Node{data: data}
	currNode := l.head
	for currNode.next != nil {
		currNode = currNode.next
	}
	currNode.next = newNode
}

func (l *LinkedList) Display() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList{}

	list.Insert(6)
	list.Insert(9)
	list.Insert(420)
	list.InsertTail(69420)

	fmt.Println("Linked list elements: ")
	list.Display()
}
