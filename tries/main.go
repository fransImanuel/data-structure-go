package main

import "fmt"

//size of alphabet a-z
const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

func (t *Trie) Insert(w string) {
	wordLength := len(w)
	// temp variable for traverse through all nodes from root
	temp := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if temp.children[charIndex] == nil {
			temp.children[charIndex] = &Node{}
		}
		temp = temp.children[charIndex]
	}
	temp.isEnd = true
}

func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	// temp variable for traverse through all nodes from root
	temp := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if temp.children[charIndex] == nil {
			return false
		}
		temp = temp.children[charIndex]
	}
	if temp.isEnd == true {
		return true
	}
	return false
}

func main() {
	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	myTrie := InitTrie()

	for _, v := range toAdd {
		myTrie.Insert(v)
	}

	fmt.Println(myTrie.Search("wizard"))
	fmt.Println(myTrie.Search("oregon"))
}
