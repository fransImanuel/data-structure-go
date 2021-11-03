package main_tutorial

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(i int) {
	h.array = append(h.array, i)
	h.HeapifyUp(len(h.array) - 1)
}

//heapify from bottom to top
func (h *MaxHeap) HeapifyUp(index int) {
	for h.array[h.parent(index)] < h.array[index] {
		h.swap(h.parent(index), index)
		index = h.parent(index)
	}
}

//Extract returns the largest key , and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0 ")
		return -1
	}

	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.HeapifyDown(0)

	return extracted
}

//heapify from top to bottom
func (h *MaxHeap) HeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := h.left(index), h.right(index)
	childToCompare := 0

	//loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { //when right child is larger
			childToCompare = r
		}

		//compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = h.left(index), h.right(index)
		} else {
			return
		}
	}
}

//get parent index
func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

//get left child index
func (h *MaxHeap) left(i int) int {
	return 2*i + 1
}

//get right child index
func (h *MaxHeap) right(i int) int {
	return 2*i + 2
}

//swap
func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}

	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
