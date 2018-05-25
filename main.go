package main

import (
	"bytes"
	"fmt"
)

func main() {
	tree := NewMerkleTrie()
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

func NewMerkleTrie() *MerkleTrie {
	return &MerkleTrie{
		root:  newNode(0, nil),
		empty: true,
	}
}

type MerkleTrie struct {
	root  *node
	empty bool
}

func (t *MerkleTrie) Add(hash []byte) {
	if t.empty {
		t.root.value = &hash
		t.empty = false
	} else {
		t.root.add(hash)
	}
}

func (t *MerkleTrie) Print() {
	traverse(t.root, "")
}

func traverse(n *node, prefix string) {
	if n.value != nil {
		fmt.Printf("%s %8b\n", prefix, *n.value)
		return
	}
	if n.left != nil {
		traverse(n.left, prefix+"0")
	}
	if n.right != nil {
		traverse(n.right, prefix+"1")
	}
}

func newNode(level byte, value *[]byte) *node {
	return &node{
		level:  level,
		number: level / 8,
		bit:    byte(1 << (7 - (level % 8))),
		value:  value,
	}
}

type node struct {
	level, bit, number byte
	left, right        *node
	value              *[]byte
}

func (n *node) add(hash []byte) {
	if n.value != nil && bytes.Equal(*n.value, hash) {
		return
	}

	var left bool
	if n.bit&hash[n.number] == 0 {
		left = true
	}

	if left {
		if n.left == nil {
			n.left = newNode(n.level+1, &hash)
		} else {
			n.left.add(hash)
		}
	} else {
		if n.right == nil {
			n.right = newNode(n.level+1, &hash)
		} else {
			n.right.add(hash)
		}
	}

	if n.value != nil {
		value := *n.value
		n.value = nil
		n.add(value)
	}
}
