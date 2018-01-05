package main

import "fmt"

func main() {
	var s string

	fmt.Scanf("%s", &s)

	for i := 0; i < len(s); i += 10 {
		if i+10 > len(s) {
			fmt.Printf("%s\n", s[i:len(s)])
		} else {
			fmt.Printf("%s\n", s[i:i+10])
		}
	}
}
