package list

import "fmt"

type list struct {
	key  int
	prev *list
	next *list
}

var data = []int{3, 4, 5, 6, 8, 910}

func RunList() {
	head := &list{
		0,
		nil,
		nil,
	}

	// hydrating the list.
	for _, d := range data {
		insert(head, d)
	}

	deleteF(head, 8)

	iterate(head)

	fmt.Println(search(head, 61))
}

func insert(head *list, key int) {
	newList := &list{
		key,
		head,
		head.next,
	}
	if head.next != nil {
		head.next.prev = newList
	}
	head.next = newList
}

func search(head *list, key int) *list {
	item := head.next

	for item != nil && item.key != key {
		item = item.next
	}

	return item
}

func deleteF(head *list, key int) {
	toDelete := search(head, key)

	if toDelete == nil {
		return
	}

	toDelete.prev.next = toDelete.next
	if toDelete.next != nil {
		toDelete.next.prev = toDelete.prev
	}
}

func iterate(head *list) {
	item := head.next
	for item != nil {
		fmt.Printf("%d, ", item.key)
		item = item.next
	}
	fmt.Println()
}
