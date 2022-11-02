package main

import "fmt"

func main() {

	type person struct {
		name string
		age  int
		pet  string
	}

	bob := person{
		name: "Bob",
		age:  25,
	}
	bob.pet = "Jake"
	fmt.Println(bob.name)
	fmt.Println(bob)

	// anonymous struct
	var anon struct {
		name string
		age  int
		pet  string
	}
	anon.name = "Wololo"
	anon.age = 50
	anon.pet = "doguinho"

	pet := struct {
		name string
		kind string
	}{
		name: "Jake",
		kind: "dog",
	}
	fmt.Println(pet)

	// converting one struct to another (everything must be the same to compile)
	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}
	g = f
	fmt.Println(f == g)
}
