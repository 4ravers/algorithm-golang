package main

import "fmt"

func main() {
	const five, three int = 5, 3
	var n int
	fiveCnt := 0
	fmt.Scanf("%d", &n)
	fiveCnt = (n / five)
	for (n-five*fiveCnt)%three != 0 {
		if fiveCnt <= 0 {
			break
		}
		fiveCnt--
	}
	if (n-five*fiveCnt)%three == 0 {
		fmt.Printf("%d\n", fiveCnt+(n-five*fiveCnt)/three)
	} else {
		fmt.Printf("%d\n", -1)
	}

}
