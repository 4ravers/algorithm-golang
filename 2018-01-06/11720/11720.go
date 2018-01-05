package main

import "fmt"

func main() {
	var d int
	var n string
	sum := 0

	fmt.Scanf("%d", &d)
	fmt.Scanf("%s", &n)

	for i := 0; i < d; i++ {
		sum += int(n[i] - '0')
	}

	fmt.Printf("%d\n", sum)

}
