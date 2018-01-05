package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d", &a)
	for i := a; i >= 1; i-- {
		fmt.Printf("%d\n", i)
	}
}
