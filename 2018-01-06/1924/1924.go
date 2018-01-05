package main

import "fmt"

func main() {
	var m, d int

	fmt.Scanf("%d %d", &m, &d)

	for i := 1; i < m; i++ {
		if i == 2 {
			d += 28
		} else if i < 8 {
			d += i%2 + 30
		} else {
			d += 31 - i%2
		}
	}

	switch d % 7 {
	case 1:
		fmt.Println("MON")
	case 2:
		fmt.Println("TUE")
	case 3:
		fmt.Println("WED")
	case 4:
		fmt.Println("THU")
	case 5:
		fmt.Println("FRI")
	case 6:
		fmt.Println("SAT")
	default:
		fmt.Println("SUN")
	}

}
