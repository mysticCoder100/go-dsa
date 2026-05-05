package tree

import "fmt"

type heap struct {
	data []int
}

func (h *heap) last() int {
	return h.data[len(h.data)-1]
}
func (h *heap) root() int {
	return h.data[0]
}

func (h *heap) rightChildIndex(index int) int {
	return (index * 2) + 2
}

func (h *heap) leftChildIndex(index int) int {
	return (index * 2) + 1
}

func (h *heap) parentIndex(index int) int {
	return (index - 1) / 2
}

func (h *heap) insert(value int)   {
	h.data = append(h.data, value)
	newIndex := len(h.data) - 1

	for (newIndex > 0) && (h.data[newIndex] > h.data[h.parentIndex(newIndex)]) {
		h.data[newIndex], h.data[h.parentIndex(newIndex)] = h.data[h.parentIndex(newIndex)], h.data[newIndex]
		newIndex = h.parentIndex(newIndex)
	}
}

func (h *heap) delete() {
	if len(h.data) == 0 {
		return
	}

	h.data[0] = h.data[len(h.data)-1]

	h.data = append(h.data[:len(h.data)-1])

	if len(h.data) == 0 {
		return
	}

	tricklingNode := 0

	for h.hasGreaterChildren(tricklingNode) {
		greaterChild := h.calculateLargerChild(tricklingNode)

		h.data[tricklingNode], h.data[greaterChild] = h.data[greaterChild], h.data[tricklingNode]

		tricklingNode = greaterChild
	}

}

func (h *heap) calculateLargerChild(i int) int {
	leftChildIndex := h.leftChildIndex(i)
	rightChildIndex, hasRightChild := h.rightChildIndex(i), len(h.data) > h.rightChildIndex(i)

	if !hasRightChild {
		return leftChildIndex
	}

	if h.data[rightChildIndex] > h.data[leftChildIndex] {
		return rightChildIndex
	}else {
		return leftChildIndex
	}
}

func (h *heap) hasGreaterChildren(i int) bool {
	leftChildIndex, hasLefChild := h.leftChildIndex(i), len(h.data) > h.leftChildIndex(i)
	rightChildIndex, hasRightChild := h.rightChildIndex(i), len(h.data) > h.rightChildIndex(i)

	return (hasLefChild && (h.data[leftChildIndex] > h.data[i])) ||
		(hasRightChild && (h.data[rightChildIndex] > h.data[i]))
}

func RunHeap() {
	heap := &heap{data: make([]int, 0)}

	heap.insert(3)
	heap.insert(10)
	heap.insert(5)
	heap.insert(18)

	fmt.Println(heap)
}