package trie

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"
)

func NewMerkleTrie() *MerkleTrie {
	return &MerkleTrie{
		root:  newNode(0, nil, nil),
		empty: true,
	}
}

type MerkleTrie struct {
	root  *node
	empty bool
}

func (t *MerkleTrie) Add(key, value []byte) {
	if t.empty {
		t.root.key = key
		t.root.value = value
		t.empty = false
	} else {
		t.root.add(key, value)
	}
}

func (t *MerkleTrie) Hash() []byte {
	if t.empty {
		h := sha256.New()
		return h.Sum(nil)
	}
	return hash(t.root)
}

func (t *MerkleTrie) MaxDepth() byte {
	if t.empty {
		return 0
	}

	var r func(n *node) byte
	r = func(n *node) byte {
		if n.key != nil {
			return n.level
		}

		var left, right byte
		if n.left != nil {
			left = r(n.left)
		}
		if n.right != nil {
			right = r(n.right)
		}
		if left > right {
			return left
		}
		return right
	}
	return r(t.root)
}

func hash(n *node) []byte {
	if n.key != nil {
		h := sha256.New()
		return h.Sum(n.value)
	}
	if n.left == nil {
		return hash(n.right)
	}
	if n.right == nil {
		return hash(n.left)
	}

	left := hash(n.left)
	right := hash(n.right)
	h := sha256.New()
	h.Write(left)
	h.Write(right)
	return h.Sum(nil)
}

func (t *MerkleTrie) Print() {
	traversePrint(t.root, "")
}

func newNode(level byte, key, value []byte) *node {
	return &node{
		level:  level,
		number: level / 8,
		bit:    byte(1 << (7 - (level % 8))),
		key:    key,
		value:  value,
	}
}

type node struct {
	level, bit, number byte
	left, right        *node
	key, value         []byte
}

func (n *node) add(key, value []byte) {
	if n.key != nil && bytes.Equal(n.key, key) {
		n.value = value
		return
	}

	var left bool
	if n.bit&key[n.number] == 0 {
		left = true
	}

	if left {
		if n.left == nil {
			n.left = newNode(n.level+1, key, value)
		} else {
			n.left.add(key, value)
		}
	} else {
		if n.right == nil {
			n.right = newNode(n.level+1, key, value)
		} else {
			n.right.add(key, value)
		}
	}

	if n.key != nil {
		key, value := n.key, n.value
		n.key = nil
		n.value = nil
		n.add(key, value)
	}
}

func traversePrint(n *node, prefix string) {
	if n.key != nil {
		fmt.Printf("%s %s\n", prefix, formatBinary(n.key))
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
