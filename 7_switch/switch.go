package main

import "fmt"

func main() {

	//normal switch
	// i := 2
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// case 4:
	// 	fmt.Println("four")
	// case 5:
	// 	fmt.Println("five")
	// default:
	// 	fmt.Println("other")
	// }

	// multiple condition switch
	// switch time.Now().Weekday(){
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it is weekend")
	// default:
	// 	fmt.Println("it is workday")
	// }

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("it is int")
		case string:
			fmt.Println("it is string")
		case bool:
			fmt.Println("it is boolean")
		default:
			fmt.Println("other", t)
		}
	}
	whoAmI(true)
}