package detMatrix

import (
	"fmt"
	"math"
	"os"
)

func endApp() {
	message := recover()
	fmt.Println(message)
}

func DetMatrix() {
	var matrix_1 [10][10]float64
	var matrix_2 [10][10]float64
	var matrix_hasil [10][10]float64
	var matrix [10][10]float64
	var row, col, row_2, col_2 int
	pilih := 0
	fmt.Println("1. penjumlahan : ")
	fmt.Println("2. pengurangan : ")
	fmt.Println("3. perkalian : ")
	fmt.Println("4. determinant : ")

	fmt.Println("masukan jenis operasi : ")
	fmt.Scanln(&pilih)

	if pilih == 1 {

		// matrix_1
		fmt.Println("masukan jumlah baris matrix 1 : ")
		fmt.Scanln(&row)

		fmt.Println("masukan jumlah kolom 1 : ")
		fmt.Scanln(&col)

		for i := 0; i < row; i++ {
			for j := 0; j < col; j++ {
				fmt.Printf("masukan nilai (%d,%d) : ", (i + 1), (j + 1))
				fmt.Scanln(&matrix_1[i][j])
			}
		}

		// matrix_2
		fmt.Println("masukan jumlah baris matrix 2 : ")
		fmt.Scanln(&row_2)

		fmt.Println("masukan jumlah kolom 2 : ")
		fmt.Scanln(&col_2)

		// checking the rules

		if row != row_2 || col != col_2 {
			fmt.Println("ordo tidak sama tidak bisa dioperasikan")
			os.Exit(3)
		}

		// matrix_2
		for i := 0; i < row_2; i++ {
			for j := 0; j < col_2; j++ {
				fmt.Printf("masukan nilai (%d,%d) : ", (i + 1), (j + 1))
				fmt.Scanln(&matrix_2[i][j])
			}
		}

		fmt.Println(" hasil penjumlahan matrix : \n")
		for i := 0; i < row; i++ {
			fmt.Print("|\t")
			for j := 0; j < col; j++ {
				matrix_hasil[i][j] = matrix_1[i][j] + matrix_2[i][j]
				fmt.Print(matrix_hasil[i][j], "\t")
			}
			fmt.Print("|\n")
		}
	} else if pilih == 2 {
		// matrix_1
		fmt.Println("masukan jumlah baris matrix 1 : ")
		fmt.Scanln(&row)

		fmt.Println("masukan jumlah kolom 1 : ")
		fmt.Scanln(&col)

		for i := 0; i < row; i++ {
			for j := 0; j < col; j++ {
				fmt.Printf("masukan nilai (%d,%d) : ", (i + 1), (j + 1))
				fmt.Scanln(&matrix_1[i][j])
			}
		}

		// matrix_2
		fmt.Println("masukan jumlah baris matrix 2 : ")
		fmt.Scanln(&row_2)

		fmt.Println("masukan jumlah kolom 2 : ")
		fmt.Scanln(&col_2)

		// checking the rules
		if row != row_2 || col != col_2 {
			fmt.Println("ordo tidak sama tidak bisa dioperasikan")
			os.Exit(3)
		}

		// matrix_2
		for i := 0; i < row_2; i++ {
			for j := 0; j < col_2; j++ {
				fmt.Printf("masukan nilai (%d,%d) : ", (i + 1), (j + 1))
				fmt.Scanln(&matrix_2[i][j])
			}
		}

		fmt.Println(" hasil matrix : \n")
		for i := 0; i < row; i++ {
			fmt.Print("|\t")
			for j := 0; j < col; j++ {
				matrix_hasil[i][j] = matrix_1[i][j] - matrix_2[i][j]
				fmt.Print(matrix_hasil[i][j], "\t")
			}
			fmt.Print("|\n")
		}
	} else if pilih == 3 {

		// matrix_1

		fmt.Println("masukan jumlah baris matrix 1 : ")
		fmt.Scanln(&row)

		fmt.Println("masukan jumlah kolom matrix 1 : ")
		fmt.Scanln(&col)

		for i := 0; i < row; i++ {
			for j := 0; j < col; j++ {
				fmt.Printf("masukan nilai (%d,%d) : ", (i + 1), (j + 1))
				fmt.Scanln(&matrix_1[i][j])
			}
		}

		// matrix_2

		fmt.Println("masukan jumlah baris matrix 2 : ")
		fmt.Scanln(&row_2)

		fmt.Println("masukan jumlah kolom matrix 2 : ")
		fmt.Scanln(&col_2)

		if col != row_2 {
			fmt.Println("ordo tidak sesuai")
			os.Exit(3)
		}

		for i := 0; i < row_2; i++ {
			for j := 0; j < col_2; j++ {
				fmt.Printf("masukan nilai (%d,%d) : ", (i + 1), (j + 1))
				fmt.Scanln(&matrix_2[i][j])
			}
		}

		for i := 0; i < row; i++ {
			for j := 0; j < col_2; j++ {
				for k := 0; k < row_2; k++ {
					matrix_hasil[i][j] += matrix_1[i][k] * matrix_2[k][j]
				}

			}
		}

		// print the result
		fmt.Println(" hasil perkalian matrix : \n")
		for i := 0; i < row; i++ {
			fmt.Print("|\t")
			for j := 0; j < col_2; j++ {
				fmt.Print(matrix_hasil[i][j], "\t")
			}
			fmt.Print("|\n")
		}

	} else if pilih == 4 {

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

		detMat := detMatrixProcess(matrix, row, col)

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
		fmt.Printf("ordo : (%d,%d) \n", row, col)
		fmt.Printf("determinant : %f", detMat)
	} else {
		fmt.Println("pilihanmu tidak ada di menu")
	}

}

func detMatrixProcess(matrix [10][10]float64, row int, col int) float64 {

	if row == 2 && col == 2 {
		var detMat float64 = matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
		return detMat

	}
	lastCol := col - 1
	rowProc := row
	for h := 1; h <= (col - 1); h++ {
		for i := 0; i < (rowProc - 1); i++ {
			lastColEx := lastCol
			hasilBagi := 0.0
			for k := 0; k <= lastCol; k++ {
				var a float64
				if i+1 == (row - 1) {
					a = matrix[row-1][lastColEx]
				} else {
					a = matrix[i+1][lastColEx]
				}

				defer endApp()

				if lastColEx == lastCol {
					hasilBagi = matrix[i][lastColEx] / a
				}

				matrix[i][lastColEx] = matrix[i][lastColEx] - a*hasilBagi

				if math.IsNaN(matrix[i][lastColEx]) {
					matrix[i][lastColEx] = 0
				}
				if math.IsInf(matrix[i][lastColEx], +1) {
					matrix[i][lastColEx] = 1
				}

				lastColEx--

			}
		}

		rowProc--
		lastCol--
	}

	// matrix determinant
	var detMat float64 = 1.0
	for i := 0; i < row; i++ {
		detMat *= matrix[i][i]
	}

	return detMat

}

// func det(data [][]float64, n int) float64 {

// 	result := 0.0
// 	if len(data) == 1 {
// 		result = data[0][0]
// 		return result
// 	}
// 	var elemlen int
// 	var arrlen int
// 	for i := 0; i < len(data); i++ {
// 		//
// 		arrlen = n - 1
// 		elemlen = len(data[0]) - 1
// 		temp := make([][]float64, arrlen)

// 		for d := range temp {
// 			temp[d] = make([]float64, elemlen)
// 		}
// 		//
// 		for j := 1; j < len(data); j++ {
// 			for k := 0; k < len(data[0]); k++ {
// 				if k < i {
// 					temp[j-1][k] = data[j][k]
// 				} else if k > i {
// 					temp[j-1][k-1] = data[j][k]
// 				}
// 			}
// 		}
// 		n--
// 		fmt.Println(i)
// 		result += data[0][i] * math.Pow(-1, float64(i)) * det(temp, n)
// 	}
// 	return result

// }
