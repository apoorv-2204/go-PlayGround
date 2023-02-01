package main

import (
	"fmt"
	"time"
)

func print_values(i int, l int) {
	for i < l {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf(" %d", i)
		i++
	}
}

func print_large(i int, m int) {
	for i < m {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf(" %d", i)
		i++
	}
}

func routine_generator(i int, l int) {
	for i < l {
		go print_values(i, i+1)
		i += 1
	}
}

func main() {

	// go print_values(1, 100)
	// go print_large(1001, 2000)
	fmt.Println("The big bang theory")
	// print_values(1
	go routine_generator(1, 100000)
	fmt.Scanln()
}
