package main

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 { return false } 
    
    for index, row := range matrix {
        
        if target >= row[0] {
            
            for _, value := range matrix[index] {
                
                if value > value {
                    break
                }
                
                if target == value {
                    return true
                }
            }
        }
    }
    
    return false
}