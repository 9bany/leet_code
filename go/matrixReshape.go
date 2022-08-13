package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if canReshape(mat, r, c) {
		return reshape(mat, r, c)
	}    
	return mat
}


func canReshape(nums [][]int, r, c int) bool  {
	row := len(nums)
	colume := len(nums[0])
	if row*colume == r*c {
		return true
	}
	return false
}

func reshape(nums [][]int, r, c int) [][]int {
	newShape := make([][]int, r)
	
	for index := range newShape {
		newShape[index] = make([]int, c)		
	}

	rowIndex, columeIndex := 0, 0
	
	for _ , row := range nums {
		for _ , col := range row {
			if columeIndex == c { 
				columeIndex = 0
				rowIndex++
			}
			newShape[rowIndex][columeIndex] = col
			columeIndex++
		}
	}

	return newShape
}

