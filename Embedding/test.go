package main

import (
	"fmt"
	"reflect"
)

// type Animal struct {
// 	Name   string
// 	Origin string
// }

// type Bird struct {
// 	Animal
// 	SpeedKPH float32
// 	CanFly   bool
// }

type Animal struct {
	Name   string `required max: "100"`
	Origin string
}

func main() {
	// b := Bird{}
	// b.Name = "Emu"
	// b.Origin = "Australia"
	// b.SpeedKPH = 48
	// b.CanFly = true
	// fmt.Println(b)

	// b := Bird{
	// 	Animal:   Animal{Name: "Emu", Origin: "Australia"},
	// 	SpeedKPH: 48,
	// 	CanFly:   true,
	// }
	// fmt.Println(b)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
