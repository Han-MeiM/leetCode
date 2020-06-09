package main

import "fmt"

func setTop(top []int, i int, j int) {
	sign := top[j]
	for k, v := range top {
		if v == sign {
			top[k] = top[i]
		}
	}
}

func equationsPossible(equations []string) bool {
	top := make([]int, 26)
	for i := 0; i < 26; i++ {
		top[i] = i
	}

	for _, v := range equations {
		if v[1] == '=' {
			tmp1 := int(v[0] - 'a')
			tmp2 := int(v[3] - 'a')
			setTop(top, tmp1, tmp2)
		}
	}

	fmt.Println(top)
	for _, v := range equations {
		if v[1] == '!' {
			tmp1 := int(v[0] - 'a')
			tmp2 := int(v[3] - 'a')
			if top[tmp1] == top[tmp2] {
				return false
			}
		}
	}

	return true
}

func main() {
	arr := []string{"a==b", "c==b", "c==d", "d!=r", "x!=z"}

	equationsPossible(arr)
}
