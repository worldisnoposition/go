package leetcode

import "fmt"

func searchMatrix(matrix [][]int,target int) bool {
	for i,x := range matrix {
        for j,_ := range x {
            if matrix[i][j] == target {
                return true
            }
        }
    }
	return false
}

func Entrance() {
	matrix := [][]int{{1,2},{3,4}}
	target := 4
	a := fmt.Sprintf("%t",searchMatrix(matrix, target))
	fmt.Printf(a)
}