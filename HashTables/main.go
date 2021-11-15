package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	Array [ArraySize]*Bucket
}

type Bucket struct {
	head *Node
}

type Node struct {
	key  string
	next *Node
}

func (h *HashTable) Insert(key string) {
	index := Hash(key)
	h.Array[index].Insert(key)
}

func (h *HashTable) Search(key string) {
	index := Hash(key)
	h.Array[index].Search(key)
}

func (h *HashTable) Delete(key string) {
	index := Hash(key)
	h.Array[index].Delete(key)
}

func (b *Bucket) Insert(key string) {
	newNode := &Node{key: key}
	if b.head == nil {
		b.head = newNode
	} else {
		newNode.next = b.head
		b.head = newNode
	}
}

func (b *Bucket) Search(key string) bool {
	temp := b.head
	for temp != nil {
		if temp.key == key {
			return true
		}
		temp = temp.next
	}
	return false
}

func (b *Bucket) Delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
	}

	temp := b.head
	for temp.next != nil {
		if temp.next.key == key {
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}
}

func Hash(key string) int {
	sum := 0
	for _, k := range key {
		sum += int(k)
	}
	return sum % ArraySize
}

func InitHashTable() *HashTable {
	init := &HashTable{}
	for i := 0; i < ArraySize; i++ {
		init.Array[i] = &Bucket{}
	}
	return init
}

func main() {

	list := []string{
		"ERIC", "KENNY", "KYLE", "STAN", "RANDY", "BUTTERS", "TOKEN",
	}

	testBucket := &Bucket{}

	for _, v := range list {
		testBucket.Insert(v)
	}

	fmt.Println(testBucket)
}
