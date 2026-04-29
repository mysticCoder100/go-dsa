package list

import (
	"fmt"
)

// creating the Node

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func initList() *linkedList {
	l := &linkedList{}
	return l
}

func (l *linkedList) read(index int)(int,error){
	currentNode := l.head
	currentIndex := 0

	for currentNode != nil {
		if currentIndex == index {
			return currentNode.data, nil
		}
		currentNode = currentNode.next
		currentIndex++
	}
	return 0, fmt.Errorf("index %d out of range", index)
}

func (l *linkedList) insertAt(index int, value int)error{
	item := &node{data: value}

	if index == 0 {
		item.next = l.head
		l.head = item

		return nil;
	}
	currentNode := l.head
	currentIndex := 0

	for currentIndex < index-1 {
		if currentNode == nil {
			return fmt.Errorf("index %d out of range", index)
		}
		currentNode = currentNode.next
		currentIndex++
	}

	if currentNode == nil {
		return fmt.Errorf("index %d out of range", index)
	}

	item.next = currentNode.next
	currentNode.next = item
	return nil
}

func (l *linkedList) deleteAt(index int)error{

	if index == 0 {
		l.head = l.head.next

		return nil;
	}
	currentNode := l.head
	currentIndex := 0

	for currentIndex < index-1 {
		if currentNode == nil {
			return fmt.Errorf("index %d out of range", index)
		}
		currentNode = currentNode.next
		currentIndex++
	}

	if currentNode == nil {
		return fmt.Errorf("index %d out of range", index)
	}

	currentNode.next = currentNode.next.next
	return nil
}

func (l *linkedList) search(value int)(int,error){
	currentNode := l.head
	currentIndex := 0

	for currentNode != nil {
		if currentNode.data == value {
			return currentIndex, nil
		}
		currentNode = currentNode.next
		currentIndex++
	}

	return 0, fmt.Errorf("value %d not found", value)
}


func Run() {
	node1 := &node{data: 3};
	node2 := &node{data: 4};
	node3 := &node{data: 45};
	node4 := &node{data: 5};

	node1.next = node2;
	node2.next = node3;
	node3.next = node4;

	l := initList();
	l.head = node1;

	//item, _ := l.read(2)

	//index,err := l.search(45)

	err := l.deleteAt(1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(l.read(1));
}
