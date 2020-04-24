// https://www.hackerrank.com/challenges/grid-challenge/problem

package main

import (
	"fmt"
	"sort"
)

func gridChallenge(grid []string) string {
	var matrix = make([][]rune, len(grid))

	for i := 0; i < len(grid); i++ {
		matrix[i] = make([]rune, len(grid[0]))

		for j := 0; j < len(grid[0]); j++ {
			matrix[i][j] = rune(grid[i][j])
		}
	}

	for i := 0; i < len(matrix); i++ {
		sort.Slice(matrix[i], func(j, k int) bool {
			return matrix[i][j] < matrix[i][k]
		})
	}

	for i := 0; i < len(matrix[0]); i++ {
		for j := 1; j < len(matrix); j++ {
			if matrix[j][i] < matrix[j-1][i] {
				return "NO"
			}
		}
	}

	return "YES"
}

func main() {
	fmt.Println(gridChallenge([]string{"ebacd", "fghij", "olmkn", "trpqs", "xywuv"}))
	fmt.Println(gridChallenge([]string{"a"}))
}
