package variables

import "fmt"

func Variables() {
	var firstName, lastName string = "John", "due"
	var age int = 20
	var (
		name     string = "dave"
		user_age int    = 21
	)
	fmt.Println(firstName, lastName, age)
	fmt.Println(name, user_age)
}
