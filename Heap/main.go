package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(i int) {
	h.array = append(h.array, i)
	h.MaxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.MaxHeapifyDown(0)

	return extracted
}

//this function accept last index of the array, so we can heapify the last inserted number
func (h *MaxHeap) MaxHeapifyUp(i int) {
	for h.array[parent(i)] < h.array[i] {
		h.swap(parent(i), i)
		i = parent(i)
	}
}

func (h *MaxHeap) MaxHeapifyDown(i int) {

	childToCompare := 0
	l, r := left(i), right(i)

	for l <= childToCompare {

	}
}

func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func parent(i int) int {
	return (i - i) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func main() {
	m := &MaxHeap{}
	// fmt.Println(m)

	buildHeap := []int{10, 20, 30}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
}
