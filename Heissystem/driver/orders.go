package driver

//import "fmt"

func Orders_make_state_matrix(N_elevators int, M_floors int) [][]int {
	
	State_matrix := make([][]int, N_elevators)
	for i := range State_matrix {
		State_matrix[i] = make([]int, M_floors*3 + 3) // evt M_floors*3 - 2 + 3
		for j := range State_matrix[i] {
			State_matrix[i][j] = 0
		}
	}
	return State_matrix
}
