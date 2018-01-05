package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d", &a)
	for i := 1; i <= a; i++ {
		fmt.Printf("%d\n", i)
	}
}
