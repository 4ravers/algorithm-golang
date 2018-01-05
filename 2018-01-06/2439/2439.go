package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d", &a)
	for i := 1; i <= a; i++ {
		for j := 1; j <= a-i; j++ {
			fmt.Printf(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
