package main

import "fmt"

// error and func values are returned generally

// func add(a int, b int) int {
// 	return a + b
// }

func getLanguages()(string, string, string) {
	return "golang", "python", "java"
}

func main() {
	// result := add(5, 10)
	// fmt.Println(getLanguages())
	// lang1, lang2, lang3 := getLanguages()
	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)
	// fmt.Println(lang1, lang2, lang3)
	// fmt.Println(`The sum is`, result)
}