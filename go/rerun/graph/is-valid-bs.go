package graph

func isValidBST(root *TreeNode) bool {
	valid := func(treeMin, curr, treeMax int) bool {
		if curr > treeMin && curr < treeMax {
			return true
		}
		return false
	}
	var dfs func(r *TreeNode, treeMin, treeMax int)
	flag := true
	dfs = func(r *TreeNode, treeMin, treeMax int) {
		if r == nil {
			return
		} else if !valid(treeMin, r.Val, treeMax) {
			flag = false
		}
		dfs(r.Left, treeMin, r.Val)
		dfs(r.Right, r.Val, treeMax)
	}
	dfs(root, -9999999999, 99999999999)
	return flag
}
