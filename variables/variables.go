package variables

import "fmt"

func Variables() {
	// var name type = expression
	var firstName, lastName string = "John", "due"
	var age int = 20
	var (
		name     string = "dave"
		user_age int    = 21
	)
	fmt.Println(firstName, lastName, age)
	fmt.Println(name, user_age)

	// name := expression
	content_in := "go"
	fmt.Println(content_in)

	// Pointers in short variable decleration
	x := 1
	p := &x
	fmt.Println(p)
	fmt.Println(*p)

	// Pointers
	var y int = 2
	var point *int = &y
	fmt.Println(point)
	fmt.Println(*point)

	// Type decleration
	type fahrenheit int
	type celsius int

	var f fahrenheit = 32
	var c celsius = 0

	fmt.Println(f, c)

	c = celsius((f - 32) * 5 / 9)
	fmt.Println(f, c)
}
