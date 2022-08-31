package main

func pacificAtlantic(matrix [][]int) [][]int {
	if matrix == nil || (matrix != nil && len(matrix) < 1) {
		return nil
	}
	rst := [][]int{}
	m, n := len(matrix), len(matrix[0])
	visited_P := make([][]bool, m)
	visited_A := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited_P[i] = make([]bool, n)
		visited_A[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		dfs(matrix, visited_P, i, 0)
		dfs(matrix, visited_A, i, n-1)
	}
	for i := 0; i < n; i++ {
		dfs(matrix, visited_P, 0, i)
		dfs(matrix, visited_A, m-1, i)
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if visited_P[r][c] && visited_A[r][c] {
				rst = append(rst, []int{r, c})
			}
		}
	}

	return rst
}

func dfs(matrix [][]int, visited [][]bool, r, c int) {

	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	m, n := len(matrix), len(matrix[0])
	visited[r][c] = true
	for _, d := range dir {
		R, C := r+d[0], c+d[1]
		if 0 <= R && 0 <= C && R < m && C < n && !visited[R][C] && matrix[r][c] <= matrix[R][C] {
			dfs(matrix, visited, R, C)

		}
	}
}