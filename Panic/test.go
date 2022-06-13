package main

import (
	"fmt"
	"log"
)

func main() {
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	// fmt.Println("start")
	// panic("something bad happened")
	// fmt.Println("end")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Println("start")
	// defer fmt.Println("this was deferred")
	// panic("something bad happened")
	// fmt.Println("end")

	// fmt.Println("start")
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		log.Println("Error:", r)
	// 	}
	// }()
	// panic("something bad happened")
	// fmt.Println("end")

	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error:", r)
			panic(r)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}
