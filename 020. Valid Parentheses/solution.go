package main

import "fmt"

func isValid(s string) bool {
	match := map[string]string{
		"{" : "}",
		"[" : "]",
		"(" : ")",
	}

	var stack []string

	for _, v := range s {
		v_str := string(v)
		if v_str == "{" || v_str == "[" || v_str == "(" {
			stack = append(stack, v_str)
		} else if len(stack) == 0 {
			return false
		} else if match[stack[len(stack) - 1]] != v_str {
			return false
		} else {
			stack = append(stack[:len(stack)-1], stack[len(stack):]...)
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	str := "[](){}"
	res := isValid(str)
	fmt.Println(res)
}
