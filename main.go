package main

import (
	"fmt"
	"log"
)

// Babylonian method. It starts with an initial guess and then updates it repeatedly until it reaches a desired level of accuracy.
func squareRoot(number float64, precision int) float64 {
	var result float64 = 1.0
	for i := 0; i < precision; i++ {

		result = 0.5 * (result + number/result)
	}
	return result
}

// PrintRectangle prints symbol in x*y times in x rows and y columns
func PrintRectangle(x int, y int, symbol string) {
	// count nb of rows, row loop
	// break row, row 0 ->
	//            row 1 ->
	// formulate formula

	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			fmt.Print(symbol)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func PrintRectangleBoundary(breadth int, length int, sym string) {

	// for perimeter print symbol
	for row := 0; row < breadth; row++ {
		if row == 0 || row == breadth-1 {
			// if first and last row print symbol length nb of times
			for j := 0; j < length; j++ {
				fmt.Print(sym)
				fmt.Print(" ")
			}
		} else {
			// if not first and not last row print symbol 2 times only at start and end of row
			// and print length -2 times space

			// print on first and last column
			fmt.Print(sym)
			fmt.Print(" ")
			for j := 0; j < length-2; j++ {
				fmt.Print("  ")

			}

			fmt.Print(sym)
			fmt.Print(" ")
		}
		fmt.Println()
	}

}

func PrintHalfPyramid(size int, symbol string) {
	// for each row print the symbol nth+1 times
	// observation and formula
	// nth row print symbol n+1 times
	for row := 0; row < size; row++ {
		for col := 0; col < row+1; col++ {
			fmt.Print(symbol)
		}
		fmt.Println()
	}
}

// +++++
// ++++
// +++
// ++
// +
func PrintInversePryamid(max_rows int, symbol string) {
	// for each nth row print max_rows - row nb

	for row_nb := 0; row_nb < max_rows; row_nb++ {
		for col := 0; col < max_rows-row_nb; col++ {
			fmt.Print(symbol)
		}
		fmt.Println()
	}

}

// 1  row_nb 0
// 1 2  row_nb 1
// 1 2 3  row_nb 2
// 1 2 3 4  row_nb 3
// how many cols to print for each row? A: row_nb + 1
// what to print? A: number 1 to row_nb + 1
func HalfNumberPryamid(size int) {
	// obsercation pattern deducation and formulation
	for row_nb := 0; row_nb < size; row_nb++ {
		// numbers are counted upto row_nb +1
		// for 0th col 1 , for nth col print n +1
		for col := 0; col < row_nb+1; col++ {
			print(col + 1)
			print(" ")
		}
		println()
	}
}

// give size n =5
//
// 1 2 3 4 5   row_nb =0  print cols n
// 1 2 3 4     row_nb =1  print cols n- row_nb
// 1 2 3       row_nb =2  print cols n- row_nb
// 1 2         row_nb =3  print cols n- row_nb
// 1          row_nb =4  print cols n- row_nb
// how many cols to print for each row n-row_nb
// what to print. 1 to n-row_nb
func ReverseHalfNumberPryamid(size int) {
	for row_nb := 0; row_nb < size; row_nb++ {
		for col_nb := 0; col_nb < size-row_nb; col_nb++ {
			print(col_nb + 1)
			print(" ")
		}
		println()
	}
}

func main() {
	// PrintRectangle(5, 30, "*")
	// PrintRectangleBoundary(7, 100, "*")
	// if _, err := fmt.Scan(&i, &j, &k); err != nil {
	// 	log.Print("  Scan for i, j & k failed, due to ", err)
	// 	return
	// }
	// PrintHalfPyramid(size, str1)
	// PrintInversePryamid(size, str1)
	// HalfNumberPryamid(size)
	var str1 string
	var size int
	_, err := fmt.Scanln(&size, &str1)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println()
	ReverseHalfNumberPryamid(size)
	
}
