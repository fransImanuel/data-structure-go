package main_trial2

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) Insert(key int) {
	//if there is no node yet
	if n.key == 0 {
		n.key = key
	}
	if n == nil {
		n = &Node{key: key}
	}
	if n.key > key {
		//go left
		if n.left == nil {
			n.left = &Node{key: key}
		} else {
			n.left.Insert(key)
		}
	} else if n.key < key {
		//go right
		if n.right == nil {
			n.right = &Node{key: key}
		} else {
			n.right.Insert(key)
		}
	}
}

func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}

	if n.key < key {
		//go right
		return n.right.Search(key)

	} else if n.key > key {
		//go left
		return n.left.Search(key)
	}

	return true
}

func main() {
	tree := &Node{}

	tree.Insert(100)
	tree.Insert(50)
	tree.Insert(200)
	tree.Insert(13)
	tree.Insert(201)
	tree.Insert(106)
	tree.Insert(64)
	tree.Insert(27)
	tree.Insert(96)
	tree.Insert(101)
	tree.Insert(170)
	tree.Insert(203)
	fmt.Println(tree)

	fmt.Println(tree.Search(501))

}
