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
	//ceritanya kan pengen nge loop sampe ke node terakhir buat inisialisasi tail
	for temp.next != nil {
		temp = temp.next
	}
	l.tail = temp
	l.length++
}

func (l *LinkedList) InsertPushBack(input int) {
	n := &node{
		data: input,
	}
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.length++
}

func (l *LinkedList) PopFront() {
	l.head = l.head.next
}

func (l *LinkedList) PopBack() {
	temp := l.head
	for temp.next != l.tail {
		temp = temp.next
	}
	l.tail = temp
	l.tail.next = nil
}

func (l *LinkedList) PrintAll() {
	temp := l.head
	for temp != nil {
		fmt.Printf("%d --> ", temp.data)
		temp = temp.next
	}
	fmt.Printf("\n")
}

func main() {
	list := &LinkedList{}
	list.InsertPushFront(3)
	list.InsertPushFront(4)
	list.PrintAll()

	list.InsertPushBack(1)
	list.PrintAll()

	list.PopFront()
	list.PrintAll()

	list.InsertPushFront(5)
	list.InsertPushFront(6)
	list.PrintAll()

	// list.PopBack()
	list.PopBack()
	list.PrintAll()
}
