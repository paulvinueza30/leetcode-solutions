package graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func createNode(val int) *Node {
	return &Node{
		val,
		make([]*Node, 0),
	}
}

func cloneGraph(n *Node) *Node {
	if n == nil {
		return nil
	}
	adj := make(map[*Node]*Node)
	q := make([]*Node, 0)
	q = append(q, n)

	for len(q) > 0 {
		curr := q[0]
		if _, ok := adj[curr]; !ok {
			adj[curr] = createNode(curr.Val)
		}
		for _, neigh := range curr.Neighbors {
			if _, ok := adj[neigh]; !ok {
				adj[neigh] = createNode(neigh.Val)
				q = append(q, neigh)
			}
			adj[curr].Neighbors = append(adj[curr].Neighbors, adj[neigh])
		}

		q = q[1:]
	}
	return adj[n]
}
