package graph

type coord struct {
	r int
	c int
	t int
}

const (
	freshFruit  = 1
	rottenFruit = 2
)

func orangesRotting(grid [][]int) int {
	freshCount := 0
	q := make([]coord, 0)
	// init fresh count
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == rottenFruit {
				q = append(q, coord{r, c, 0})
			} else if grid[r][c] == freshFruit {
				freshCount++
			}
		}
	}
	t := 0
	addToQ := func(r, c, nextT int) {
		if r >= len(grid) || r < 0 || c >= len(grid[0]) || c < 0 || grid[r][c] != freshFruit {
			return
		}
		grid[r][c] = rottenFruit
		freshCount--
		t = nextT
		q = append(q, coord{r, c, nextT})
	}
	for len(q) > 0 {
		cd := q[0]
		q = q[1:]
		grid[cd.r][cd.c] = rottenFruit
		addToQ(cd.r+1, cd.c, cd.t+1)
		addToQ(cd.r-1, cd.c, cd.t+1)
		addToQ(cd.r, cd.c+1, cd.t+1)
		addToQ(cd.r, cd.c-1, cd.t+1)
	}
	if freshCount != 0 {
		return -1
	}
	return t
}
