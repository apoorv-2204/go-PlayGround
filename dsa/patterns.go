package dsa

import "fmt"

func Loop() {
	// for i := 1; i < 11; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 1; i <= 5; i += 2 {
	// 	fmt.Println(i)
	// }

	// for z := 100; z > 0; z /= 2 {
	// 	fmt.Println(z)
	// }
	// var z int
	// for z = 5; z >= 5 && z <= 10; z += 1 {
	// 	fmt.Println(z)
	// }

	// for a := 1; a < 10; a++ {
	// 	fmt.Println(a)
	// }

	// for {
	// 	fmt.Println("s")
	// }
	// var str1 string

	// if fmt.Scanln(&str1) {

	// }

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

//          *
//         * *
//        * * *
//       * * * *
//      * * * * *
//     * * * * * *
//    * * * * * * *
//   * * * * * * * *
//  * * * * * * * * *
// * * * * * * * * * *

// nb of rows =size
// print space in each before row = size- row_nb-1
// print symbol row_nb+1 and space
func PrintFullPryamid(size int, symbol string) {
	for row_nb := 0; row_nb < size; row_nb++ {
		for space := 0; space < size-row_nb-1; space++ {
			fmt.Print(" ")
		}
		for sym := 0; sym < row_nb+1; sym++ {
			// print(sym)
			fmt.Print(symbol)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

// * * * * * * * * * *
//  * * * * * * * * *
//   * * * * * * * *
//    * * * * * * *
//     * * * * * *
//      * * * * *
//       * * * *
//        * * *
//         * *
//---------_*
// nb_of rows = size
// first row nb of symbol size
// how many times to print symbol? : size- row_nb  size, size-1, size-2,..size-row_nb
// how many times to print space? print spce = row_nb
//

func PrintInvertedFullPryamid(size int, symbol string) {
	for row_nb := 0; row_nb < size; row_nb++ {
		for space := 0; space < row_nb; space++ {
			fmt.Print(" ")
		}
		for sym := 0; sym < size-row_nb; sym++ {
			// print(sym)
			fmt.Print(symbol)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func PrintDiamond(size int, symbol string) {
	PrintFullPryamid(size, symbol)
	PrintInvertedFullPryamid(size, symbol)
}

// |	* row 0 -> 1
// |   * *  1 ->   3
// |  *   *  2      5
// | *     * 3      7
// |*       * 4     9
// chars to print => 2*row_nb +1
// nb symbol to print => if row_nb+1
func PrintHollowPryamid(size int, symbol string) {
	for row_nb := 0; row_nb < size; row_nb++ {
		for space := 0; space < size-row_nb-1; space++ {
			fmt.Print(" ")
		}
		for cols := 0; cols < 2*row_nb+1; cols++ {
			if cols == 0 {
				// first col or first charchter to print/
				fmt.Print(symbol)
			} else if cols == 2*row_nb {
				// last col or last col to print , print symbol
				fmt.Print(symbol)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// 0|*       *      0 space  2 star   9 chars  internal_space 7
// 1| *     *       1        2		  7						  5
// 2|  *   *        2 		 2		  5						  3
// 3|   * *         3        2		  3						  1
// 4|  	 *          4 		 1		  1						  0/-1
// nb of rows to print  =>  size
// nb of chars to print => 2*size + 1 - 2*row_nb
// nb of space to print => rown_nb
// when to print star => first star => when col ==0 start of loop after space loop
//
//	last star at last columng how to get last col=> 2*size
func PrintInverseHollowPryamid(size int, symbol string) {
	for row_nb := 0; row_nb < size; row_nb++ {
		for space := 0; space < row_nb; space++ {
			fmt.Print(" ")
		}
		for cols := 0; cols < 2*size-2*row_nb-1; cols++ {
			if cols == 0 {
				// first col or first charchter to print/
				fmt.Print(symbol)
			} else if cols == 2*size-2*row_nb-2 {
				// last col or last col to print , print symbol
				fmt.Print(symbol)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// +++++++++++++++++
// ++++++++ ++++++++
// +++++++   +++++++
// ++++++     ++++++
// +++++       +++++
// ++++         ++++
// +++           +++
// ++             ++
// +               +
func PrintNotOfHollowPryamid(size int, symbol string) {
	for row_nb := 0; row_nb < size; row_nb++ {
		for space := 0; space < size-row_nb-1; space++ {
			fmt.Print(symbol)
		}
		for cols := 0; cols < 2*row_nb+1; cols++ {
			if cols == 0 {
				// first col or first charchter to print/
				fmt.Print(symbol)
			} else if cols == 2*row_nb {
				// last col or last col to print , print symbol
				fmt.Print(symbol)
			} else {
				fmt.Print(" ")
			}
		}
		for space := 0; space < size-row_nb-1; space++ {
			fmt.Print(symbol)
		}
		fmt.Println()
	}
}

// +               +
// ++             ++
// +++           +++
// ++++         ++++
// +++++       +++++
// ++++++     ++++++
// +++++++   +++++++
// ++++++++ ++++++++
// +++++++++++++++++
// we can combine multiple triangles also
func PrintNotOfInverseHollowPryamid(size int, symbol string) {
	for row_nb := 0; row_nb < size; row_nb++ {
		for space := 0; space < row_nb; space++ {
			fmt.Print(symbol)
		}
		for cols := 0; cols < 2*size-2*row_nb-1; cols++ {
			if cols == 0 {
				// first col or first charchter to print/
				fmt.Print(symbol)
			} else if cols == 2*size-2*row_nb-2 {
				// last col or last col to print , print symbol
				fmt.Print(symbol)
			} else {
				fmt.Print(" ")
			}
		}
		for space := 0; space < row_nb; space++ {
			fmt.Print(symbol)
		}
		fmt.Println()
	}
}

//	nb of rows = 4
//
// 1
// 2*2
// 3*3*3
// 4*4*4*4
// 5*5*5*5*5
// 6*6*6*6*6*6
// print row_nb+1 row_nb+1 times
func NumberMountain(size int) {

	// not inverse
	for row_nb := 0; row_nb < size; row_nb++ {
		for col_nb := 0; col_nb < row_nb+1; col_nb++ {
			fmt.Print(row_nb + 1)
			if !(col_nb == row_nb) {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
	// print size-row_nb   size-row_nb times
	// inverse

	for row_nb := 0; row_nb < size; row_nb++ {
		for col_nb := 0; col_nb < size-row_nb; col_nb++ {
			fmt.Print(size - row_nb)
			// if not last column print
			// nb of iteration -1
			if !(col_nb == size-row_nb-1) {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
	// 1
	// 2 * 2
	// 3 * 3 * 3
	// 4 * 4 * 4 * 4
	// 5 * 5 * 5 * 5 * 5
	// 6 * 6 * 6 * 6 * 6 * 6
	// 7 * 7 * 7 * 7 * 7 * 7 * 7
	// 7 * 7 * 7 * 7 * 7 * 7 * 7
	// 6 * 6 * 6 * 6 * 6 * 6
	// 5 * 5 * 5 * 5 * 5
	// 4 * 4 * 4 * 4
	// 3 * 3 * 3
	// 2 * 2
	// 1
}

// |----1
// |---232
// |--34543
// |-4567654
// |567898765
// nb_rows= size
// print space : size-row_nb-1
// print
//   - forward symbol times= row_nb, start forward= row_nb+1,lessthan = row_nb + size -1
//     print middle row_nb +1
//   - Backward symbol times=  row_nb, start backward from= row_nb
func PrintNumericFullPryamid(size int) {
	for row_nb := 0; row_nb < size; row_nb++ {
		for space := 0; space < size-row_nb-1; space++ {
			fmt.Print(" ")
		}
		for forward := row_nb + 1; forward < row_nb+1; forward++ {
			fmt.Print(forward)
			fmt.Print(" ")
		}
		fmt.Print(row_nb + 1)
		for backward := row_nb; backward < row_nb; backward++ {
			fmt.Print(backward)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
