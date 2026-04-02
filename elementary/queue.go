package elementary

import "fmt"

type Queue struct {
	head   int
	tail   int
	length int
	items  []int
}

func RunQueue() {
	queue := MakeQueue(6)
	queue.Enqueue(4)
	queue.Enqueue(1)
	queue.Enqueue(3)
	queue.Dequeue()
	fmt.Println(queue)
}

func MakeQueue(size int) *Queue {
	return &Queue{
		head:   0,
		tail:   0,
		length: size,
		items:  make([]int, size),
	}
}

func (q *Queue) Enqueue(i int) {
	q.items[q.tail] = i
	if q.tail == q.length {
		q.tail = 0
	} else {
		q.tail += 1
	}
}

func (q *Queue) Dequeue() int {
	item := q.items[q.head]
	if q.head == q.length {
		q.head = 0
	} else {
		q.head += 1
	}
	return item
}
