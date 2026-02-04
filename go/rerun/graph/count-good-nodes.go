package graph

const (
	Max = 999999
	Min = -999999
)

func goodNodes(root *TreeNode) int {
	var helper func(r *TreeNode, maximum int) int
	helper = func(r *TreeNode, mm int) int {
		if r == nil {
			return 0
		}
		newMax := max(r.Val, mm)
		if newMax == Max {
			newMax = r.Val
		}
		if r.Val < mm && mm != Max {
			return helper(r.Left, newMax) + helper(r.Right, newMax)
		}
		return 1 + helper(r.Left, newMax) + helper(r.Right, newMax)
	}
	return helper(root, Max)
}
