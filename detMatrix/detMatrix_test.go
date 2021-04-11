package detMatrix

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDetMatrixProcess(t *testing.T) {

	// testing 2 * 2 matrix
	var matrix [10][10]float32

	matrix[0][0] = 2
	matrix[0][1] = 2
	matrix[1][0] = 3
	matrix[1][1] = 4

	detMat := detMatrixProcess(&matrix, 2, 2)
	assert.Equal(t, detMat, 2.0, "they should be equal")

	// testing n * n matrix
	var matrix2 [10][10]float32

	matrix2[0][0] = 1
	matrix2[0][1] = 2
	matrix2[0][2] = 3
	matrix2[0][3] = 6

	matrix2[1][0] = 2
	matrix2[1][1] = 2
	matrix2[1][2] = 2
	matrix2[1][3] = 4

	matrix2[2][0] = 2
	matrix2[2][1] = 3
	matrix2[2][2] = 2
	matrix2[2][3] = 2

	matrix2[3][0] = 1
	matrix2[3][1] = 2
	matrix2[3][2] = 4
	matrix2[3][3] = 4

	detMat = detMatrixProcess(&matrix2, 4, 4)
	assert.Equal(t, detMat, -28.0, "they should be equal")
}
