package main

import (
	"fmt"
	"maps"
)

func main(){
	// m := make(map[string]string)
	// m["name"] = "vansh"
	// m["area"] = "punjab"
	// m["lang"] = "golang"

	// fmt.Println(m["name"])
	// fmt.Println(m["area"])
	// fmt.Println(m["lang"])

	// m := map[string]int{"price":30,"phone":6564564}
	// fmt.Println(m)

	// check if element exist in map
	// value, ok := m["price"]
	// fmt.Println(value)
	// if ok{
	// 	fmt.Println("it exists")
	// }else{
	// 	fmt.Println("it exists")
	// }

	m1 := map[string]int{"price":30,"phone":6564564}
	m2 := map[string]int{"price":30,"phone":2323}

	fmt.Println(maps.Equal(m1,m2))
	

}