package main

import "fmt"

func CountIf(f func(string) bool, tab []string) int {
	count := 0

	for i := 0; i < len(tab); i++ {
		if f(tab[i]) {
			count++
		}
	}
	return count
}

func main() {
	tab := []string{"Bonjour", "le", "monde", "!"}
	f := func(s string) bool {
		return len(s) > 3
	}
	result := CountIf(f, tab)
	fmt.Println(result)
}
