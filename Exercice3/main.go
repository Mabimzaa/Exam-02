package main

import "fmt"

func SplitWhiteSpaces(s string) []string {
	var result []string
	start := -1

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if start != -1 {
				result = append(result, s[start:i])
				start = -1
			}
		} else {
			if start == -1 {
				start = i
			}
		}
	}
	if start != -1 {
		result = append(result, s[start:])
	}
	return result
}

func main() {
	fmt.Println(SplitWhiteSpaces("Bonjour je m'appelle Flavien"))
	fmt.Println(SplitWhiteSpaces("  Bonjour   je m'appelle   Flavien  "))
}
