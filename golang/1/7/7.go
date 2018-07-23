package problems

// Write an algorithm such that if an element in an MxN
// matrix is 0, its entire row and
// column is set to 0.
func destroyRowsCols(matrix [][]int) {
	// use map as set to save the rows/cols to be deleted
	rowsDelete := make(map[int]bool)
	colsDelete := make(map[int]bool)

	// calculate the dimensions
	rowCount := len(matrix)
	colCount := len(matrix[0])

	// Iterate over each element
	for row := 0; row < rowCount; row++ {
		for col := 0; col < colCount; col++ {
			// Add the row and col to the deletion set if
			// cur element is zero
			if matrix[row][col] == 0 {
				rowsDelete[row] = true
				colsDelete[col] = true
			}
		}
	}

	// destroy rows and cols
	for row := 0; row < rowCount; row++ {
		for col := 0; col < colCount; col++ {
			if rowsDelete[row] || colsDelete[col] {
				matrix[row][col] = 0
			}
		}
	}

}
