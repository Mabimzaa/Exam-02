package main

import "fmt"

func IsSorted(f func(x int) int, a []int) bool {
	cro := true
	dec := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i]) > f(a[i+1]) {
			cro = false
		}
		if f(a[i]) < f(a[i+1]) {
			dec = false
		}
	}
	return cro || dec
}

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	f := func(x int) int {
		return x
	}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
