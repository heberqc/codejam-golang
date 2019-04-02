package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scanf("%d", &T)
	T = T + 1
	for i := 1; i < T; i++ {
		var sMax int
		var audience string
		var output int = 0
		fmt.Scanf("%d %s", &sMax, &audience)
		standing := CharToNum(audience[0])
		for shyness := 1; shyness <= sMax; {
			if standing < shyness {
				standing++
				output++
			} else if standing >= shyness {
				standing = standing + CharToNum(audience[shyness])
				shyness++
			}
		}
		fmt.Printf("Case #%d: %d\n", i, output)
	}
}

func CharToNum(r byte) int {
	return int(r) - '0'
}
