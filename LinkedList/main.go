package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	length int
	head   *node
	tail   *node
}

func (l *LinkedList) InsertPushFront(input int) {
	n := &node{
		data: input,
	}
	if l.head == nil {
		l.head = n
	} else {
		n.next = l.head
		l.head = n
	}

	temp := l.head
	for temp != nil {
		temp = temp.next
	}
	l.tail = temp
	l.length++

	fmt.Printf("\nHead : %v", l.head)
	fmt.Printf("\nTail : %v", l.tail)
}

func (l *LinkedList) InsertPushBack(input int) {
	n := &node{
		data: input,
	}
	if l.head == nil {
		l.head = n
	} else {
		fmt.Println(l.tail.data)
		l.tail.next = n
		l.tail = n
	}
	l.length++
}

func (l *LinkedList) PrintAll() {
	temp := l.head
	for temp != nil {
		fmt.Printf("%d --> ", temp.data)
		temp = temp.next
	}
}

func main() {
	list := &LinkedList{}
	list.InsertPushFront(1)
	list.InsertPushFront(2)
	// fmt.Println(list)
	// fmt.Println(list.head.data)
	// fmt.Println(list.head.next.data)
	list.PrintAll()
	// list.InsertPushBack(3)
	// list.PrintAll()
	// fmt.Println(list)
}
