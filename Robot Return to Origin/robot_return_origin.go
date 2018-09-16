package main

import "fmt"

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, v := range moves {
		switch string(v) {
		case "R":
			x++
		case "L":
			x--
		case "U":
			y--
		case "D":
			y++
		}
	}
	return x == 0 && y == 0
}

func main() {
	str := "UD"
	res := judgeCircle(str)
	fmt.Println(res)
}
