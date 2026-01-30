package graph

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var res int
	var helper func(r *TreeNode)
	helper = func(r *TreeNode) {
		if r == nil {
			return
		}
		helper(r.Left)
		k -= 1
		if k == 0 {
			res = r.Val
		}
		helper(r.Right)
	}

	helper(root)
	return res
}
