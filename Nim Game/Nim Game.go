package main

import "fmt"

func canWinNim(n int) bool {
	if n % 4 == 0 {
		return false
	} else {
		return true
	}
}

func main() {
	res := canWinNim(4)
	fmt.Println(res)
}
