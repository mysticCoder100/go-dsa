package tree

type trieNode struct {
	children map[rune]*trieNode
}

type trie struct {
	root *trieNode
}

func (t *trie) search(word string) *trieNode {
	currNode := t.root

	for _, char := range word {

		if node,ok := currNode.children[char]; ok {
			currNode = node
		}else {
			return nil
		}
	}

	return currNode
}

func (t* trie) collectWords() *[]string {
	var result []string
	return t.collectAllWords(t.root, "", &result)
}

func (t* trie) collectAllWords (node *trieNode, word string, words *[]string) *[]string {
	currNode := node
	for index, value := range currNode.children {
		if index == '*' {
			*words = append(*words, word);
		} else {
			t.collectAllWords(value, word + string(index), words)
		}
	}

	return words
}

func (t *trie) insert(word string) {
	currNode := t.root

	for _, char := range word {
		if node,ok := currNode.children[char]; ok {
			currNode = node
		} else {
			child := &trieNode{make(map[rune]*trieNode)}
			currNode.children[char] = child
			currNode = child
		}
	}

	currNode.children['*'] = nil
}