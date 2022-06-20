package main

import (
	"fmt"
    "math"
	// "math/rand"
    "time"
)

// print Hello World! message

func printMessage() {
	fmt.Println("Hello World!")
}

// print name & age

func printName() {
	fmt.Println("My name is Touhidul Shawan\nI am 24 years old")
}

// print numbers from 1 to 10

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i)
		fmt.Println()
	}
}

// raise 2 to the power of 11

func numPower(){
    output := math.Pow(2,11)
    fmt.Println(output)
}

// generate random number  
//does not complete yet

// func randomNumber(){
//     randNumber := rand.Int31()
//     fmt.Println(randNumber)
// }

// print current date

func printDate(){
    currentTime := time.Now()
    fmt.Println(currentTime)
}

func main() {
	printMessage()
	printName()
	printNumbers()
    numPower()
    printDate()
    // randomNumber()
}
