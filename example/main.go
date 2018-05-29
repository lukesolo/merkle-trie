package main

import (
	"github.com/lukesolo/trie"
)

func main() {
	tree := trie.NewMerkleTrie()

	b1 := []byte{1 << 7}
	b2 := []byte{1 << 6}
	b3 := []byte{(1 << 7) + (1 << 6)}
	tree.Add(b1)
	tree.Add(b2)
	tree.Add(b3)
	tree.Add([]byte{1, 5})
	tree.Add([]byte{1, 4})
	tree.Add([]byte{0, 0})
	tree.Add([]byte{255})

	tree.Print()
}
