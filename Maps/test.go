package main

import "fmt"

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12801989,
		"Illinois":     12741080,
		"Ohio":         11614373,
	}
	// statePopulations["Ohio"] = 8938389
	// fmt.Println(statePopulations)
	// delete(statePopulations, "Ohio")
	// fmt.Println(statePopulations)

	// // _, ok := statePopulations["Ohio"]
	// // fmt.Println(ok)
	// fmt.Println(len(statePopulations))

	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)
}
