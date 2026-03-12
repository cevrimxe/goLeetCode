func maxScore(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    const INF = int(1e18)

    minVal := make([][]int, m)
    for i := range minVal {
        minVal[i] = make([]int, n)
        for j := range minVal[i] {
            minVal[i][j] = INF
        }
    }

    ans := -INF

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i > 0 {
                minVal[i][j] = min(minVal[i][j], minVal[i-1][j])
            }
            if j > 0 {
                minVal[i][j] = min(minVal[i][j], minVal[i][j-1])
            }

            if minVal[i][j] != INF {
                ans = max(ans, grid[i][j]-minVal[i][j])
            }

            minVal[i][j] = min(minVal[i][j], grid[i][j])
        }
    }

    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
