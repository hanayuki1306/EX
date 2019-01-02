// https://leetcode.com/problems/maximum-depth-of-binary-tree/

package main

import (
	"fmt"
	"io"
	"os"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxDepth(root *BinaryNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.left), maxDepth(root.right))

}
func main() {

	tree := &BinaryTree{}

	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10).
		insert(95).
		insert(98).
		insert(99)
	f := maxDepth(tree.root)
	print(os.Stdout, tree.root, 0, 'M')
	fmt.Println(f)
}

// 			100
// 		-20
// 	-50 	-15
// -60				50
// 			15		60
// 		-10		55		85
// 					5		95
// 								98
// 									99

// package main
// import (
// 	"fmt"
// 	"math/rand"
// )

// // A Tree is a binary tree with integer values.
// type Tree struct {
// 	Left  *Tree
// 	Value int
// 	Right *Tree
// }

// // Walk traverses a tree depth-first,
// // sending each Value on a channel.
// func Walk(t *Tree, ch chan int) {
// 	if t == nil {
// 		return
// 	}
// 	Walk(t.Left, ch)
// 	ch <- t.Value
// 	Walk(t.Right, ch)
// }

// // Walker launches Walk in a new goroutine,
// // and returns a read-only channel of values.
// func Walker(t *Tree) <-chan int {
// 	ch := make(chan int)
// 	go func() {
// 		Walk(t, ch)
// 		close(ch)
// 	}()
// 	return ch
// }

// // Compare reads values from two Walkers
// // that run simultaneously, and returns true
// // if t1 and t2 have the same contents.
// func Compare(t1, t2 *Tree) bool {
// 	c1, c2 := Walker(t1), Walker(t2)
// 	for {
// 		v1, ok1 := <-c1
// 		v2, ok2 := <-c2
// 		if !ok1 || !ok2 {
// 			return ok1 == ok2
// 		}
// 		if v1 != v2 {
// 			break
// 		}
// 	}
// 	return false
// }

// // New returns a new, random binary tree
// // holding the values 1k, 2k, ..., nk.
// func New(n, k int) *Tree {
// 	var t *Tree
// 	for _, v := range rand.Perm(n) {
// 		t = insert(t, (1+v)*k)
// 	}
// 	return t
// }

// func insert(t *Tree, v int) *Tree {
// 	if t == nil {
// 		return &Tree{nil, v, nil}
// 	}
// 	if v < t.Value {
// 		t.Left = insert(t.Left, v)
// 		return t
// 	}
// 	t.Right = insert(t.Right, v)
// 	return t
// }

// func main() {
// 	t1 := New(100, 1)
// 	fmt.Println(Compare(t1, New(100, 1)), "Same Contents")
// 	fmt.Println(Compare(t1, New(99, 1)), "Differing Sizes")
// 	fmt.Println(Compare(t1, New(100, 2)), "Differing Values")
// 	fmt.Println(Compare(t1, New(101, 2)), "Dissimilar")
// }

// nghĩa là t có 1 cái cây và t mới import xong cái cây đó nhưng ko rõ nó
