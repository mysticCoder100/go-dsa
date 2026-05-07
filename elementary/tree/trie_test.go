package tree

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func createTree() *trie {
	return &trie{
		root: &trieNode{
			children: make(map[rune]*trieNode),
		},
	}
}

func populateTree(tree *trie) {
	path,err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fullPath := filepath.Join(path, "words.txt")

	data,er := os.ReadFile(fullPath)

	if er != nil {
		panic(er)
	}
	content := string(data)

	for _, line := range strings.Split(content, "\n") {
		tree.insert(line)
	}
}


func TestCanInsertToTrie(t *testing.T) {
	tr := createTree()

	tr.insert("cat")
	tr.insert("can")

	// 1. Verify 'c' exists
	cNode, cOk := tr.root.children['c']
	if !cOk {
		t.Fatal("Expected node 'c' to exist")
	}

	// 2. Verify 'a' exists under 'c'
	aNode, aOk := cNode.children['a']
	if !aOk {
		t.Fatal("Expected node 'a' to exist under 'c'")
	}

	// 3. Verify 't' and 'n' branch off from 'a'
	if _, tOk := aNode.children['t']; !tOk {
		t.Error("Missing 't' in 'cat'")
	}
	if _, nOk := aNode.children['n']; !nOk {
		t.Error("Missing 'n' in 'can'")
	}
}

func TestGetWords(t *testing.T) {
	tr := createTree()
	populateTree(tr)

	result := tr.collectWords()
	fmt.Println((*result)[0])
}