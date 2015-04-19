package driver

import "fmt"

func Orders_make_matrix(N_floors int) [][]int {
	
	Matrix := make([][]int, N_floors)
	for i := range Matrix {
		Matrix[i] = make([]int, 13)
		for j := range Matrix[i] {
			Matrix[i][j] = 0
		}
	}
	fmt.Println(Matrix)
	return Matrix
}
