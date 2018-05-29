package main

import (
	"crypto/rand"
	"fmt"

	"github.com/lukesolo/trie"
)

func main() {
	depth()

	tree := trie.NewMerkleTrie()

	b1 := []byte{1 << 7}
	b2 := []byte{1 << 6}
	b3 := []byte{(1 << 7) + (1 << 6)}
	tree.Add(b1, b1)
	tree.Add(b2, b2)
	tree.Add(b3, b3)
	tree.Add([]byte{1, 5}, nil)
	tree.Add([]byte{1, 4}, nil)
	tree.Add([]byte{0, 0}, nil)
	tree.Add([]byte{255}, nil)

	tree.Print()
	fmt.Printf("%x", tree.Hash())
}

func depth() {
	tree := trie.NewMerkleTrie()

	for i := 0; i < 1000000; i++ {
		key := make([]byte, 32)
		rand.Read(key)
		tree.Add(key, key)
	}

	fmt.Println(tree.MaxDepth())
}
