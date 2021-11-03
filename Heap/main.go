package main

import "fmt"

type Heap struct {
	array []int
}

func (h *Heap) Insert(i int) {
	h.array = append(h.array, i)
	// h.MaxHeapifyTop(len(h.array) - 1)
	h.MinHeapifyTop(len(h.array) - 1)
}

func (h *Heap) Extract() int {
	l := len(h.array) - 1
	extracted := h.array[0]
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	// h.MaxHeapifyDown(0)
	h.MinHeapifyDown(0)

	return extracted
}

//Max Heapify from bottom to top
func (h *Heap) MaxHeapifyTop(i int) {
	//bedanya minheapify dan maxheapify cma terjadi di if ini
	for h.array[parent(i)] < h.array[i] {
		h.swap(parent(i), i)
		i = parent(i)
	}
}

//Min Heapify from bottom to top
func (h *Heap) MinHeapifyTop(i int) {
	//bedanya minheapify dan maxheapify cma terjadi di if ini
	for h.array[parent(i)] > h.array[i] {
		h.swap(parent(i), i)
		i = parent(i)
	}
}

//Max Heapify from top to bottom
func (h *Heap) MaxHeapifyDown(i int) {
	lastIndex := len(h.array) - 1
	childToCompare := 0
	l, r := left(i), right(i)

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		//bedanya minheapify dan maxheapify cma terjadi di if ini
		if h.array[i] < h.array[childToCompare] {
			h.swap(i, childToCompare)
			i = childToCompare
			l, r = left(i), right(i)
		} else {
			return
		}

	}
}

//Max Heapify from top to bottom
func (h *Heap) MinHeapifyDown(i int) {
	lastIndex := len(h.array) - 1
	childToCompare := 0
	l, r := left(i), right(i)

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] < h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		//bedanya minheapify dan maxheapify cma terjadi di if ini
		if h.array[i] > h.array[childToCompare] {
			h.swap(i, childToCompare)
			i = childToCompare
			l, r = left(i), right(i)
		} else {
			return
		}

	}
}

func (h *Heap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return i*2 + 1
}
func right(i int) int {
	return i*2 + 2
}

func main() {
	m := &Heap{}

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	var sortedArray []int
	for _, _ = range m.array {
		extract := m.Extract()
		fmt.Println(m)
		sortedArray = append(sortedArray, extract)
	}
	fmt.Printf("Sort Result = %d", sortedArray)
}
