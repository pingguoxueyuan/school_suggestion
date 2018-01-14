package logic
/*
import (
	"fmt"
)
*/
type Node struct {
	char rune
	childs map[rune]*Node
	Data interface{}
	deep int
	isTerm bool
}

type Trie struct {
	root *Node
	size int
}

func NewNode(char rune, deep int) *Node {
	node := &Node {
		char: char,
		childs:make(map[rune]*Node, 16),
		deep:deep,
	}
	return node
}

func NewTrie() *Trie {
	trie := &Trie {
		root: NewNode(' ', 1),
		size:1,
	}
	return trie
}

func (t *Trie) Add(key string, data interface{}) {
	var parent *Node = t.root
	allChars := []rune(key)
	for _, char := range allChars {
		node, ok := parent.childs[char]
		if !ok {
			node = NewNode(char, parent.deep+1)
			parent.childs[char] = node
		}

		parent = node
	}

	parent.Data = data
	parent.isTerm = true
}

func (t *Trie) PrefixSearch(key string, limit int) (nodes []*Node) {

	var node = t.root
	allChars := []rune(key)
	for _, char := range allChars {
		child, ok := node.childs[char]
		if !ok {
			//fmt.Printf("prefix char:%c\n", char)
			return
		}

		node = child
	}

	//fmt.Printf("prefix node:%+v\n", node)
	var queue []*Node
	queue = append(queue, node)
	for len(queue) > 0 {
		var q2 []*Node
		for _, n := range queue {
			if n.isTerm == true {
				//fmt.Printf("n is leaf node:%v", n)
				nodes = append(nodes, n)
				if len(nodes) > limit {
					return
				}
				continue
			}
			
			for _, v := range n.childs {
				q2 = append(q2, v)
			}
		}
		queue = q2
	}

	return
}