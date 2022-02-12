package main

import (
	"fmt"
)

func arrayInGo() {
	// Array in go
	var names = [10]string{"Mikle", "June", "Sam", "Cat"}
	fmt.Printf("Array : %v\n", names)
	fmt.Printf("Single value: %v\n", names[2])
	fmt.Printf("Length of array: %v\n", len(names))

	// Another way to declare array

	var fruits [20]string

	fruits[0] = "Mango"
	fruits[1] = "Apple"

	fmt.Printf("Fruit Array : %v\n", fruits)

}

// slice in array
// slice is an abstraction of array
// slice is use to create dynamic size array

func sliceInGo() {

	var persons = []string{}

	// add new elements to persons array

	persons = append(persons, "Mike")
	persons = append(persons, "Sam")

	fmt.Printf("Persons : %v\n", persons)
	fmt.Printf("Lenght of Persons array : %v\n", len(persons))

	// another way to declare slice

	superHero := []string{}

	superHero = append(superHero, "Iron Man")
	superHero = append(superHero, "Bat Man")
	superHero = append(superHero, "Hulk")

	fmt.Printf("Super Heroes: %v\n", superHero)
}

func main() {
	arrayInGo()
	sliceInGo()

}
