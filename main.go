package main

import (
	"fmt"
	"log"
)

func main() {
	arr := []int{
		16, 17, 18, 1, 1, 1, 2, 3, 4, 5, 6, 7, 1, 2,
	}

}

func TakeInput() {
	var str1 string
	var size int
	_, err := fmt.Scanln(&size, &str1)
	if err != nil {
		log.Fatal(err)
	}
}
