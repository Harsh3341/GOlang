package main

import "fmt"

// const (
// 	a = iota
// 	b
// 	c
// 	d
// )

// const (
// 	a2 = iota
// )

// const (
// 	catSpecialist = iota
// 	dogSpecialist
// 	snakeSpecialist
// )

// const (
// 	_  = iota             // ignore first value by assigning to blank identifier
// 	KB = 1 << (iota * 10) // 1 << (1 * 10)
// 	MB                    // 1 << (2 * 10)
// 	GB                    // 1 << (3 * 10)
// 	TB                    // 1 << (4 * 10)
// 	PB                    // 1 << (5 * 10)
// 	EB                    // 1 << (6 * 10)
// 	ZB                    // 1 << (7 * 10)
// 	YB                    // 1 << (8 * 10)
// )

const (
	isAdmin          = 1 << iota // 1 << 0
	isHeadquarters               // 1 << 1
	canSeeFinancials             // 1 << 2

	canSeeAfrica       // 1 << 3
	canSeeAsia         // 1 << 4
	canSeeEurope       // 1 << 5
	canSeeNorthAmerica // 1 << 6
	canSeeSouthAmerica // 1 << 7
)

func main() {
	// const myConst int = 42
	// fmt.Printf("%v, %T\n", myConst, myConst)

	// const a int = 14
	// const b string = "foo"
	// const c float32 = 3.14
	// const d bool = true
	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)
	// fmt.Printf("%v\n", c)
	// fmt.Printf("%v\n", d)

	// const a int = 14
	// fmt.Printf("%v,%T\n", a, a)

	// fmt.Printf("%v,%T\n", a, a)

	// fmt.Printf("%v,%v,%v,%v,%v", a, b, c, d, a2)

	// var specialistType int = catSpecialist
	// fmt.Printf("%v\n", specialistType == catSpecialist)

	// fileSize := 4000000000.
	// fmt.Printf("%.2fGB\n", fileSize/GB)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is admin? %v\n", roles&isAdmin == isAdmin)
	fmt.Printf("Can see financials? %v\n", roles&canSeeFinancials == canSeeFinancials)
	fmt.Printf("Can see Europe? %v\n", roles&canSeeEurope == canSeeEurope)
	fmt.Printf("Can see North America? %v\n", roles&canSeeNorthAmerica == canSeeNorthAmerica)
}
