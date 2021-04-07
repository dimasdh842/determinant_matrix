package detMatrix

import "fmt"

func endApp() {
	message := recover()
	fmt.Println(message)
}

func DetMatrix() {

	var matrix [100][100]int
	var row, col int
	fmt.Println("masukan jumlah baris : ")
	fmt.Scanln(&row)

	fmt.Println("masukan jumlah kolom : ")
	fmt.Scanln(&col)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("masukan nilai (%d,%d) : ", (i + 1), (j + 1))
			fmt.Scanln(&matrix[i][j])
		}
	}

	// detMatrixProcess(matrix, row, col)
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Println("matrix : ")
	for i := 0; i < row; i++ {
		fmt.Print("|\t")
		for j := 0; j < col; j++ {
			fmt.Print(matrix[i][j], "\t")
		}

		fmt.Print("|")
		fmt.Print("\n")
	}
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Printf("ordo : (%d,%d)", row, col)
}

func detMatrixProcess(matrix [100][100]int, row int, col int) {

	for i := 0; i < row; i++ {
		for j := 1; j < col; j++ {
			defer endApp()
			lastCol := col
			matrix[i][lastCol] = matrix[i][lastCol] - (matrix[i][lastCol] / matrix[i+1][lastCol] * matrix[i+1][lastCol])
			lastCol--
		}
	}
}
