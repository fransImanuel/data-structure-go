package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*Bucket
}

type Bucket struct {
	head *Node
}

type Node struct {
	key  string
	next *Node
}

//Init HashTable
func InitHashTable() *HashTable {
	hashT := &HashTable{}
	for i := range hashT.array {
		hashT.array[i] = &Bucket{}
	}
	return hashT
}

func (h *HashTable) Insert(key string) {
	index := Hash(key)
	h.array[index].Insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := Hash(key)
	return h.array[index].Search(key)
}

func (h *HashTable) Delete(key string) {
	index := Hash(key)
	h.array[index].Delete(key)
}

//Insert Bucket
func (b *Bucket) Insert(key string) {
	if b.Search(key) {
		fmt.Println("already exist")
		return
	}
	newNode := &Node{key: key}
	if b.head == nil {
		b.head = newNode
	} else {
		newNode.next = b.head
		b.head = newNode
	}
}

//Search Bucket
func (b *Bucket) Search(key string) bool {
	traverse := b.head
	for traverse != nil {
		if traverse.key == key {
			return true
		}
		traverse = traverse.next
	}
	return false
}

//Delete Bucket
func (b *Bucket) Delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
	}

	traverse := b.head
	for traverse.next != nil {
		if traverse.next.key == key {
			traverse.next = traverse.next.next
		}
		traverse = traverse.next
	}
}

func Hash(w string) int {
	sum := 0

	for _, i := range w {
		fmt.Println(i)
		sum += int(i)
	}

	return sum % ArraySize
}

func main() {
	// hashT := InitHashTable()
	// fmt.Println(hashT)

	list := []string{
		"ERIC", "KENNY", "KYLE", "STAN", "RANDY", "BUTTERS", "TOKEN",
	}

	testBucket := &Bucket{}

	for _, v := range list {
		testBucket.Insert(v)
	}
}
