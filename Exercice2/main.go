package main

import "fmt"

func ConcatParams(params []string) string {
	result := ""
	for _, param := range params {
		result += param
		result += "\n"
	}
	return result
}

func main() {
	params := []string{"Bonjour", "le", "monde", "4&!1"}
	result := ConcatParams(params)
	fmt.Println(result)
}
