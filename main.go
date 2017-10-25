package main

import (
	"fmt"
)

func errorCheck(num int) byte {
	mask := 0x80000000
	count := byte(0)
	for mask > 0 {
		check := num & mask
		if check == mask {
			count ^= 1
		}
		mask >>= 1
	}
	return count
}

func checkForString(s string) byte {
	count := byte(0)
	for _, char := range s {
		count ^= errorCheck(int(char))
	}
	return count
}
func main() {
	fmt.Println(checkForString("abcde")) // 1
	fmt.Println(errorCheck(867))         // 0
}
