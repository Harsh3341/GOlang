package main

import (
	"fmt"
	"math"
)

func main() {
	// if true {
	// 	fmt.Println("This is true")
	// }

	// statePopulations := map[string]int{
	// 	"California":   39250017,
	// 	"Texas":        27862596,
	// 	"Florida":      20612439,
	// 	"New York":     19745289,
	// 	"Pennsylvania": 12801989,
	// 	"Illinois":     12801539,
	// 	"Ohio":         11614373,
	// }
	// if pop, ok := statePopulations["Illinois"]; ok {
	// 	fmt.Println(pop)
	// }

	// number := 50
	// guess := 101

	// if guess < number {
	// 	fmt.Println("Too low")
	// }
	// if guess > number {
	// 	fmt.Println("Too high")
	// }
	// if guess == number {
	// 	fmt.Println("You got it!")
	// }
	// fmt.Println(number <= guess, number >= guess, number != guess)

	// if guess < 1 || returnTrue() || guess > 100 {
	// 	fmt.Println("Guess must be between 1 and 100")
	// }
	// if guess >= 1 && guess <= 100 {
	// 	if guess < number {
	// 		fmt.Println("Too low")
	// 	}
	// 	if guess > number {
	// 		fmt.Println("Too high")
	// 	}
	// 	if guess == number {
	// 		fmt.Println("You got it!")
	// 	}
	// 	fmt.Println(number <= guess, number >= guess, number != guess)

	// }

	// if guess < 1 {
	// 	fmt.Println("Guess must be greater than 1")
	// } else if guess > 100 {
	// 	fmt.Println("Guess must be less than 100")
	// } else {
	// 	if guess < number {
	// 		fmt.Println("Too low")
	// 	}
	// 	if guess > number {
	// 		fmt.Println("Too high")
	// 	}
	// 	if guess == number {
	// 		fmt.Println("You got it!")
	// 	}
	// 	fmt.Println(number <= guess, number >= guess, number != guess)
	// }

	myNum := 0.9
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are not the same")
	}

}

// func returnTrue() bool {
// 	fmt.Println("returning true")
// 	return true
// }
