package main

import "fmt"

const age = 20

func main(){
	const name = "golang"
	// name = "js"
	fmt.Println(age)
	fmt.Println(name)
	// const age = 30
	// fmt.Println(age)

	const (
		port = 3000
		host = "localhost"
	)
	fmt.Println(port,host);
}
