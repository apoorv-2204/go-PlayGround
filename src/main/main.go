package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	// "github.com/apoorv-2204/go-PlayGround/dsa"
	"github.com/apoorv-2204/go-PlayGround/naive_math"
)

func main() {

	fmt.Println(naive_math.SquareRoot(225454545454545, 1000))
	fmt.Println("hello")
	fmt.Println('G')
	var nb int = 923372036854775807
	fmt.Println(nb)
	fmt.Println(math.MaxFloat32, math.MaxFloat64)
	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	fmt.Println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)

	sum, mul, sub := calc(os.Args[1], os.Args[2])
	fmt.Println(sum, mul, sub)

}

func calc(number1 string, number2 string) (sum int, mul int, sub int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	sub = int1 - int2
	return
}

func TakeInput() {
	var str1 string
	var size int
	_, err := fmt.Scanln(&size, &str1)
	if err != nil {
		log.Fatal(err)
	}
}
