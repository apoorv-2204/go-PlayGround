package main

import (
	"go_playground/dsa"
)
// import alias "import/path"
// alias.Method()

func main() {

	// var str1 string
	// var size int
	// _, err := fmt.Scanln(&size, &str1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	dsa.PrintDiamond(5, "x")
	dsa.PrintFullPryamid(5, "x")
	dsa.PrintNotOfHollowPryamid(5, "+")

}

// binary as it splits searcg space in two
// func BinarySearch() {
// 	target: =2
// 	slice := []int{1, 2, 3, 4, 5}
// }
