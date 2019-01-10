
https://leetcode.com/problems/sum-of-left-leaves/submissions/

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil  && root.Left.Right == nil && root.Left.Left == nil{
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}