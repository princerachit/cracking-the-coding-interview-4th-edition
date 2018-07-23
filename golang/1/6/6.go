package problems

// Given an image represented by an NxN matrix, where each pixel in the image is 4
// bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
func rotateMatrix(matrix [][]int64) [][]int64 {
	N := len(matrix)
	// create a copy of the matrix so that testing becomes easy
	image := make([][]int64, N)
	for a := 0; a < N; a++ {
		image[a] = make([]int64, N)
	}
	for a := 0; a < N; a++ {
		for b := 0; b < N; b++ {
			image[a][b] = matrix[a][b]
		}
	}
	// Work layer wise. The outer most layer then one layer below it and so on
	for layer := 0; layer < N/2; layer++ {
		// The first layer is the index of first element
		first := layer
		// Last index
		last := N - layer - 1

		for i := first; i < last; i++ {
			// we will use offset to find the mapping element on next edge
			offset := i - first
			// get the top layer element
			top := image[first][i]
			// update top layer value to left layer mapping value
			// The y component would remain fix as it is the
			// left most element but the x component would change
			// according to the distance/offset
			image[first][i] = image[last-offset][first]

			// update the left layer element
			// The bottom layers x component is constant i.e. last
			image[last-offset][first] = image[last][last-offset]

			// update the bottom layer of the image
			// The y component would remain constant
			image[last][last-offset] = image[i][last]

			// update the left most layer with the value we saved earlier
			image[i][last] = top
		}
	}
	return image
}
