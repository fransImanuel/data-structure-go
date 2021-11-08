package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) Insert(i int) {
	if n.key == 0 && n.left == nil && n.right == nil {
		n.key = i
	}
	if n == nil {
		n.key = i
	}
	if n.key < i {
		//go right
		if n.right == nil {
			n.right = &Node{key: i}
		} else {
			n.right.Insert(i)
		}
	} else if n.key > i {
		//go left
		if n.left == nil {
			n.left = &Node{key: i}
		} else {
			n.left.Insert(i)
		}
	}
}

func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}
	if n.key < key {
		//go to right
		return n.right.Search(key)
	} else if n.key > key {
		//go to left
		return n.left.Search(key)
	}
	return true
}

func main() {
	tree := &Node{}
	tree.Insert(100)
	tree.Insert(200)
	tree.Insert(50)
	fmt.Println(tree)

	fmt.Print(tree.Search(50))
}
