class Node:
    def __init__(self):
        self.arr = [None] * 26
        self.isWord = False

class PrefixTree:

    def __init__(self):
        self.trie = Node()
    def getIdx(self, c: str) -> int:
        return ord(c) - 97
    def insert(self, word: str) -> None:
        curr = self.trie
        for c in word:
            idx = self.getIdx(c)
            if not curr.arr[idx]:
                curr.arr[idx] = Node() 
            curr = curr.arr[idx]
        curr.isWord = True
    def search(self, word: str) -> bool:
        curr = self.trie
        for c in word:
            idx = self.getIdx(c)
            if not curr.arr[idx]:
                return False
            curr = curr.arr[idx]
        return curr.isWord
    def startsWith(self, prefix: str) -> bool:
        curr = self.trie
        for c in prefix:
            idx = self.getIdx(c)
            if not curr.arr[idx]:
                return False
            curr = curr.arr[idx]
        return True
        