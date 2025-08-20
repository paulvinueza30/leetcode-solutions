package main

type Node2 struct {
	children [26]*Node2
	isWord   bool
}

type WordDictionary struct {
	root Node2
}

func Constructor2() WordDictionary {
	return WordDictionary{root: Node2{}}
}

func (t *WordDictionary) GetIdx(c rune) int {
	return int(c) - int('a')
}

func (t *WordDictionary) AddWord(w string) {
	curr := &t.root

	for _, c := range w {
		idx := t.GetIdx(c)
		if curr.children[idx] == nil {
			curr.children[idx] = &Node2{}
		}
		curr = curr.children[idx]
	}
	curr.isWord = true
}
func (t *WordDictionary) dfs(w string, n *Node2) bool {
	if n == nil {
		return false
	}
	if len(w) == 0 {
		return n.isWord
	}
	c := w[0]
	if c == '.' {
		for _, c := range n.children {
			if t.dfs(w[1:], c) {
				return true
			}
		}
		return false
	}
	cIdx := t.GetIdx(rune(c))
	if n.children[cIdx] == nil {
		return false
	}
	return t.dfs(w[1:], n.children[cIdx])

}
func (t *WordDictionary) Search(w string) bool {
	return t.dfs(w, &t.root)
}
