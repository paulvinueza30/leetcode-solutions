package linkedlist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func createNode(v int) *Node {
	return &Node{
		v,
		nil,
		nil,
	}
}

func getListLen(head *Node) int {
	l := 0
	for head != nil {
		l += 1
		head = head.Next
	}
	return l
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	q := make([]*Node, 0)
	v := make(map[*Node]*Node)
	q = append(q, head)

	for len(q) > 0 {
		curr := q[0]
		if _, ok := v[curr]; !ok {
			v[curr] = createNode(curr.Val)
		}
		cpy := v[curr]

		if _, ok := v[curr.Next]; !ok && curr.Next != nil {
			v[curr.Next] = createNode(curr.Next.Val)
			q = append(q, curr.Next)
		}
		if _, ok := v[curr.Random]; !ok && curr.Random != nil {
			v[curr.Random] = createNode(curr.Random.Val)
			q = append(q, curr.Random)
		}
		cpy.Next = v[curr.Next]
		cpy.Random = v[curr.Random]
		q = q[1:]
	}

	return v[head]
}
