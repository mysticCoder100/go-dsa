package tree

import (
	"testing"
)

func TestHeapInsert(t *testing.T) {
	heap := &heap{data: make([]int, 0)}

	heap.insert(3)
	heap.insert(10)
	heap.insert(5)
	heap.insert(18)

	expected := 18

	if heap.data[0] != expected {
		t.Errorf("Heap root is incorrect. Got %d, want %d", heap.data[0], expected)
	}

	expected = 10
	if heap.data[1] != expected {
		t.Errorf("Heap root's left child is incorrect. Got %d, want %d", heap.data[1], expected)
	}

	expected = 5
	if heap.data[2] != expected {
		t.Errorf("Heap root's right child is incorrect. Got %d, want %d", heap.data[1], expected)
	}
}

func TestHeapDelete(t *testing.T) {
	heap := &heap{data: make([]int, 0)}

	heap.insert(3)
	heap.insert(10)
	heap.insert(5)
	heap.insert(18)

	heap.delete()

	expected := 10
	if heap.data[0] != expected {
		t.Errorf("Heap root is incorrect. Got %d, want %d", heap.data[0], expected)
	}

	expected = 3

	if heap.data[1] != expected {
		t.Errorf("Heap root's left child is incorrect. Got %d, want %d", heap.data[1], expected)
	}

}
