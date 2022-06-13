package main

import "fmt"

func main() {
	// switch 4 {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("not one or two")
	// }

	// switch i := 2 + 6; i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("not one or two")
	// }

	// i := 10
	// switch {
	// case i <= 10:
	// 	fmt.Println("i is less than or equal to 10")
	// 	// fallthrough
	// case i <= 20:
	// 	fmt.Println("i is less than or equal to 20")
	// default:
	// 	fmt.Println("i is greater than 20")
	// }

	// var i interface{} = []int{}
	// switch i.(type) {
	// case int:
	// 	fmt.Println("i is an int")
	// case string:
	// 	fmt.Println("i is a string")
	// case float64:
	// 	fmt.Println("i is a float64")
	// case []int:
	// 	fmt.Println("i is an int slice")
	// default:
	// 	fmt.Println("i is a different type")
	// }

	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		break
		fmt.Println("This will print too")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is a different type")
	}
}
