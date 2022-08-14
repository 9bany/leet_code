package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 { return false }
	
	for i := 0; i < len(matrix); i++ {
		if target == matrix[i][0] { return true }
    
		if target > matrix[i][0] {
					
			for j := 0; j < len(matrix[i]); j++ {
				if target == matrix[i][j] { return true }
			}
		}
	}
	
	return false    
}
