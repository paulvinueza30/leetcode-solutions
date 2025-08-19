package main

type Node struct {
	children [26]*Node
	isWord   bool
}

type Trie struct {
	root Node
}

func Constructor() Trie {
	return Trie{root: Node{}}
}

func (t *Trie) GetIdx(c rune) int {
	return int(c) - int('a')
}

func (t *Trie) Insert(word string) {
	curr := &t.root

	for _, c := range word {
		idx := t.GetIdx(c)

		if curr.children[idx] == nil {
			curr.children[idx] = &Node{}
		}
		curr = curr.children[idx]
	}

	curr.isWord = true
}

func (t *Trie) Search(word string) bool {
	curr := &t.root

	for _, c := range word {
		idx := t.GetIdx(c)

		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return curr.isWord
}

func (t *Trie) Startswith(word string) bool {
	curr := &t.root

	for _, c := range word {
		idx := t.GetIdx(c)

		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return true
}
