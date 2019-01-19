// https://leetcode.com/problems/path-sum-iii/

package main

import (
	"fmt"
	"io"
	"os"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

type BinaryTree struct {
	root *TreeNode
}

func (t *BinaryTree) insert(Val int) *BinaryTree {
	if t.root == nil {
		t.root = &TreeNode{Val: Val, Left: nil, Right: nil}
	} else {
		t.root.insert(Val)
	}
	return t
}

func (n *TreeNode) insert(Val int) {
	if n == nil {
		return
	} else if Val <= n.Val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: Val, Left: nil, Right: nil}
		} else {
			n.Left.insert(Val)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Val: Val, Left: nil, Right: nil}
		} else {
			n.Right.insert(Val)
		}
	}
}

func print(w io.Writer, node *TreeNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Val)
	print(w, node.Left, ns+2, 'L')
	print(w, node.Right, ns+2, 'R')
}

////////////////////////////////
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return sumUp(root, 0, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func sumUp(root *TreeNode, pre int, sum int) int {
	if root == nil {
		return 0
	}
	current := pre + root.Val

	c := 0

	if current == sum {
		c = 1
	}
	return c + sumUp(root.Left, current, sum) + sumUp(root.Right, current, sum)
}

////////////////////
func main() {

	tree := &BinaryTree{}

	tree.insert(5).
		insert(4).
		insert(8).
		insert(11).
		insert(13).
		insert(4).
		insert(7).
		insert(2).
		// insert(null).
		// insert(null).
		insert(1).
		// insert(-10).
		// insert(95).
		// insert(98).
		insert(99)
	print(os.Stdout, tree.root, 0, 'M')
	// fmt.Println(tree)

	fmt.Println(pathSum(tree.root, 22))
}
