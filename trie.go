package trie

import (
	"bytes"
	"fmt"
	"strings"
)

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
	traversePrint(t.root, "")
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

func traversePrint(n *node, prefix string) {
	if n.value != nil {
		fmt.Printf("%s %s\n", prefix, formatBinary(*n.value))
		return
	}
	if n.left != nil {
		traversePrint(n.left, prefix+"0")
	}
	if n.right != nil {
		traversePrint(n.right, prefix+"1")
	}
}

func formatBinary(bs []byte) string {
	strs := make([]string, len(bs), len(bs))
	for i, b := range bs {
		strs[i] = fmt.Sprintf("%08b", b)
	}
	return strings.Join(strs, "")
}
