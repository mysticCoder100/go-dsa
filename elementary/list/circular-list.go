package list

import "fmt"

type cLink struct {
	head *cList
	tail *cList
}

type user struct {
	id         int
	name       string
	age        int
	department string
}

type cList struct {
	data user
	next *cList
	prev *cList
}

func makeCLink() *cLink {
	head := &cList{
		data: user{0, "", 0, ""},
		next: nil,
		prev: nil,
	}

	tail := &cList{
		data: user{0, "", 0, ""},
		next: head,
		prev: head,
	}
	head.next = tail
	head.prev = tail

	return &cLink{head, tail}
}

func (link *cLink) insertFront(new user) {
	newList := &cList{
		data: new,
		next: link.head.next,
		prev: link.head,
	}

	link.head.next.prev = newList
	link.head.next = newList
}

func (link *cLink) insertBack(new user) {
	newList := &cList{
		data: new,
		next: link.tail,
		prev: link.tail.prev,
	}

	link.tail.prev.next = newList
	link.tail.prev = newList
}

func (link *cLink) search(key int) *user {
	item := link.inSearch(key)

	if item != nil {
		return &item.data
	}

	return nil
}

func (link *cLink) inSearch(key int) *cList {
	item := link.head.next

	for item != link.tail && item.data.id != key {
		item = item.next
	}

	return item
}

func (link *cLink) display() {
	item := link.head.next

	for item != link.tail {
		fmt.Println(item.data)
		item = item.next
	}
}
func (link *cLink) popFront() user {
	item := link.head.next
	if item != link.tail {
		value := item.data
		item.prev.next = item.next
		item.next.prev = item.prev
		return value
	}
	return user{0,"", 0, ""}
}
func (link *cLink) popBack() user {
	item := link.tail.prev
	if item != link.head {
		value := item.data
		item.next.prev = item.prev
		item.prev.next = item.next
		return value
	}
	return user{0,"", 0, ""}
}

func (link *cLink) delete(key int) user {
	item := link.inSearch(key)
	if item != nil {
		value := item.data
		item.next.prev = item.prev
		item.prev.next = item.next
		return value
	}
	return user{0,"", 0, ""}
}



func RunCList() {
	list := makeCLink()
	
	data := [4]user {
		{1, "John", 12, "Human Resource"},
		{2,  "Haleemah", 20, 	"marketing"},
		{3, "Olivia", 13, "Development"},
		{4, "Jack", 15, "Marketing"},
	}

	dataTwo := [4]user {
		{5, "Bolu", 12, "Human Resource"},
		{6,  "Sayo", 20, 	"marketing"},
		{7, "Titi", 13, "Development"},
		{8, "Bolatito", 15, "Marketing"},
	}

	for _, d := range data {
		list.insertFront(d)
	}

	for _, d := range dataTwo {
		list.insertBack(d)
	}

	list.display()

	fmt.Println()
	list.popFront()
	list.popFront()
	list.popBack()

	list.delete(6)

	list.display()
}



