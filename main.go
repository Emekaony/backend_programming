package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int8
}

func GreetPerson(person *Person) {
	message := fmt.Sprintf("Hello there %s %s", person.FirstName, person.LastName)
	fmt.Println(message)
}

func main() {
	var ptr *Person

	Emeka := Person{
		FirstName: "Nnaemeka",
		LastName:  "Onyeokoro",
		Age:       24,
	}

	ptr = &Emeka

	GreetPerson(ptr)
}
