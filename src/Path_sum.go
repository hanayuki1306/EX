https://leetcode.com/problems/path-sum/submissions/

package main

import (
	"fmt"
	"io"
	"os"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	data  int
}

type BinaryTree struct {
	root *TreeNode
}

func (t *BinaryTree) insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &TreeNode{data: data, Left: nil, Right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *TreeNode) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.Left == nil {
			n.Left = &TreeNode{data: data, Left: nil, Right: nil}
		} else {
			n.Left.insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{data: data, Left: nil, Right: nil}
		} else {
			n.Right.insert(data)
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
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.Left, ns+2, 'L')
	print(w, node.Right, ns+2, 'R')
}

// copy code o day. doi root.data => root.Val la dc vi t tao treenode la data
func checkSumRecur(root *TreeNode, checkSum int, sum int) bool {
	if root == nil {
		return false
	}

	checkSum += root.data
	if root.Left == nil && root.Right == nil && checkSum == sum {
		return true
	}
	return checkSumRecur(root.Left, checkSum, sum) || checkSumRecur(root.Right, checkSum, sum)
}

func hasPathSum(root *TreeNode, sum int) bool {
	return checkSumRecur(root, 0, sum)
}

// den day
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

	fmt.Println(hasPathSum(tree.root, 22))
}
