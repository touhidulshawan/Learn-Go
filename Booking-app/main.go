package main

import "fmt"

func main() {
	fmt.Println("Welcome to our Booking App")

	conferenceName := "Love Yourself"

	fmt.Printf("Buy tickets for current conference: %v\n", conferenceName)
    fmt.Println("----------------------------------------")

	// take information from user

	fmt.Print("Enter your first name: ")
	var firstName string
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	var lastName string
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	var userEmail string
	fmt.Scan(&userEmail)

	fmt.Print("How many tickets you want to buy: ")
	var ticketToBuy int32
	fmt.Scan(&ticketToBuy)

    fmt.Println()
    fmt.Println("<-------- Confirmation Message -------->")

	fmt.Printf("Thanks %v %v for buying %v tickets. Your ticket is confirmed.\n", firstName, lastName, ticketToBuy)
	fmt.Printf("Check your mailbox [%v] for further details.\n", userEmail)

}
