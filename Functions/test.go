package main

import "fmt"

func main() {

	// sayMessage("Hello")
	// for i := 0; i < 10; i++ {
	// 	sayMessage("Hello", i)
	// }

	// greeting := "Hello"
	// name := "World"
	// sayGreeting(&greeting, &name)
	// fmt.Println(name)

	// s := sum(1, 2, 3, 4, 5)
	// fmt.Println(s)

	// d := divide(10, 2)
	// fmt.Println(d)

	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		} else {
			return a / b, nil
		}
	}
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

// func sayMessage(message string, idx int) {
// 	fmt.Println(message)
// 	fmt.Println("The value of the index is", idx)
// }

// func sayGreeting(greeting, name *string) {
// 	fmt.Println(*greeting, *name)
// 	*name = "Ted"
// 	fmt.Println(*name)
// }

// func sum(v ...int) int {
// 	fmt.Println(v)
// 	result := 0
// 	for _, value := range v {
// 		result += value
// 	}
// 	return result
// }

// func divide(a, b int) int {
// 	return a / b
// }
