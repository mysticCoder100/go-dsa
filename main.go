package main

import (
	f "fmt"

	dynamicprogramming "github.com/mysticCoder100/go-dsa/elementary/dynamic-programming"
)

func main() {
	dynamicprogramming.DpTest()
}

func anagram(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}
	var collection []string

	subAnams := anagram(s[1:])

	for _, v := range subAnams {
		for i := 0; i < len(v); i++ {
			f.Print(s[i])
		}
	}

	f.Println()

	return collection
}
