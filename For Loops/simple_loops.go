package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i, j := 0, 0; i <= 10; i, j = i+1, j+1 {
	// 	fmt.Println(i, j)
	// }

	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// 	if i%2 == 0 {
	// 		i /= 2
	// 	} else {
	// 		i = 2*i + 1
	// 	}
	// }

	// i := 0
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// i := 0
	// for {
	// 	fmt.Println(i)
	// 	i++
	// 	if i > 10 {
	// 		break
	// 	}
	// }

	// for i := 0; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// Loop:
	// 	for i := 1; i <= 3; i++ {
	// 		for j := 1; j <= 3; j++ {
	// 			fmt.Println(i * j)
	// 			if i*j >= 3 {
	// 				break Loop
	// 			}
	// 		}
	// 	}

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for k, v := range s {
		fmt.Println(k, v)

	}
}
