package tree

import "fmt"

type treeNode struct {
	data       int
	leftChild  *treeNode
	rightChild *treeNode
}

type bst struct {
	root *treeNode
}

func initTree() *bst {
	return &bst{}
}

func (b *bst) search(data int) (int, error) {
	currentNode := b.root

	for currentNode != nil {
		if currentNode.data == data {
			return data, nil
		} else if data < currentNode.data {
			currentNode = currentNode.leftChild
		} else {
			currentNode = currentNode.rightChild
		}
	}

	return 0, fmt.Errorf("%d not found", data)
}

func (b *bst) insert(data int) {
	currentNode := b.root
	newNode := &treeNode{data: data}

	if currentNode == nil {
		b.root = newNode
		return
	}

	for currentNode != nil {
		if data > currentNode.data {
			// move right
			if currentNode.rightChild == nil {
				currentNode.rightChild = newNode
				return
			}
			currentNode = currentNode.rightChild
		} else if data < currentNode.data {
			// move left
			if currentNode.leftChild == nil {
				currentNode.leftChild = newNode
				return
			}
			currentNode = currentNode.leftChild
		} else {
			return
		}
	}

}

func (b *bst) delete(data int) error {
	/*
		1. no child
		2. one child
		3. both child
	*/

	/*
		Perform search to locate the node.
	*/

	currentNode := b.root
	var parentNode *treeNode
	var isLeft bool

	for currentNode != nil {
		temp := currentNode
		if currentNode.data == data {
			break
		} else if data < currentNode.data {
			currentNode = currentNode.leftChild
			isLeft = true
		} else {
			currentNode = currentNode.rightChild
			isLeft = false
		}
		parentNode = temp
	}

	if currentNode == nil {
		return fmt.Errorf("%d not found", data)
	}

	if currentNode.leftChild == nil {
		if parentNode == nil {
			b.root = currentNode.rightChild
		} else if isLeft {
			parentNode.leftChild = currentNode.rightChild
		} else {
			parentNode.rightChild = currentNode.rightChild
		}
		return nil
	} else if currentNode.rightChild == nil {
		if parentNode == nil {
			b.root = currentNode.leftChild
		} else if isLeft {
			parentNode.leftChild = currentNode.leftChild
		} else {
			parentNode.rightChild = currentNode.leftChild
		}
		return nil
	}

	successor, parent := findSuccessor(currentNode)

	if successor == nil || parent == nil {
		return fmt.Errorf("%d not found", data)
	}

	currentNode.data = successor.data

	if parent.rightChild == successor {
		parent.rightChild = successor.rightChild
	} else {
		parent.leftChild = successor.rightChild
	}

	return nil
}

func findSuccessor(node *treeNode) (*treeNode, *treeNode) {
	if node == nil {
		return nil, nil
	}

	successor := node.rightChild
	parent := node

	for successor.leftChild != nil {
		parent = successor
		successor = successor.leftChild
	}

	return successor, parent
}

func RunBst() {
	tree := initTree()

	node1 := &treeNode{data: 25}
	node2 := &treeNode{data: 75}
	root := &treeNode{50, node1, node2}

	tree.root = root

	//_, err := tree.search(75)
	//
	//tree.insert(75)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//fmt.Println(tree.root)
	tree.insert(10)
	tree.insert(65)
	tree.insert(74)

	fmt.Println(root)
	//fmt.Println(root.rightChild)

	tree.delete(50)

	fmt.Println(root)
	fmt.Println(root.rightChild.leftChild)
}
