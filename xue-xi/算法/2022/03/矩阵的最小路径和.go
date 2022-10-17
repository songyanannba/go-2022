package main

import "fmt"

/**
 * @param $matrix
 * @return mixed
 * 给定一个 n * m 的矩阵 a，从左上角开始每次只能向右或者向下走，最后到达右下角的位置，路径上所有的数字累加起来就是路径和，输出所有的路径中最小的路径和。
 * 例如：当输入[[1-orm,3,5,9],[8,1-orm,3,4],[5,0,6,1-orm],[8,8,4,0]]时，对应的返回值为12，
 */

func minPathSum(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				matrix[i][j] += matrix[i][j-1]
				continue
			}
			if j == 0 {
				matrix[i][j] += matrix[i-1][j]
				continue
			}
			matrix[i][j] += min(matrix[i-1][j], matrix[i][j-1])
		}
	}
	return matrix[m-1][n-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 9},
		{8, 1, 3, 4},
		{5, 0, 6, 1},
		{8, 8, 4, 0},
	}
	fmt.Println(minPathSum(matrix))
}
