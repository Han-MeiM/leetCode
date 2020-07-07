package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if m == 0 || n == 0 || obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	ways := make([]int, n)

	for i := 0; i < m; i++ {
		if i == 0 {
			ways[0] = 1
		} else if obstacleGrid[i][0] == 1 || ways[0] == 0 {
			ways[0] = 0
		} else {
			ways[0] = 1
		}

		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				ways[j] = 0
			} else {
				ways[j] += ways[j-1]
			}
		}
	}

	return ways[n-1]
}

func main() {
	obstacleGrids := [][][]int{
		{{0, 0}, {1, 0}},
		{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		{{0, 0}},
		{{1, 0}},
		{{0}, {1}},
		{{0, 0}, {1, 1}, {0, 0}},
		{{0, 1, 0, 0, 0}, {1, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
	}

	obstacleGridWays := []int{1, 2, 1, 0, 0, 0, 0}

	result := make([]int, len(obstacleGrids))

	for ix, value := range obstacleGrids {
		result[ix] = uniquePathsWithObstacles(value)
	}

	fmt.Println(result)
	fmt.Println(obstacleGridWays)
}
