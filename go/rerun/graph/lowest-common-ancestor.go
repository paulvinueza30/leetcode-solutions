package graph

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var find func(r *TreeNode, t *TreeNode) bool
	find = func(r *TreeNode, t *TreeNode) bool {
		if r == nil {
			return false
		} else if r == t {
			return true
		}
		return find(r.Left, t) || find(r.Right, t)
	}
	var lca *TreeNode
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if (find(r, p) && r != p) && (find(r, q) && r != q) {
			lca = r
		}
		dfs(r.Left)
		dfs(r.Right)
	}
	dfs(root)
	if lca == nil {
		if p.Val > q.Val {
			lca = p
		} else {
			lca = q
		}
	}
	return lca
}
