package main

import "fmt"

func main() {
	fmt.Println("Welcome to our Booking App")

	conferenceName := "Love Yourself"
	var availableTicket uint32
	availableTicket = 200

	fmt.Printf("Buy tickets for current conference: %v\n", conferenceName)
	fmt.Printf("Available Tickets : %v\n", availableTicket)
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
	var ticketToBuy uint32
	fmt.Scan(&ticketToBuy)

	bookings := []string{}

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Println()
	fmt.Println("<-------- Confirmation Message -------->")

	fmt.Printf("Thanks %v %v for buying %v tickets. Your ticket is confirmed.\n", firstName, lastName, ticketToBuy)
	fmt.Printf("Check your mailbox [%v] for further details.\n", userEmail)
	fmt.Printf("Bookings : %v\n", bookings)

	availableTicket = availableTicket - ticketToBuy

	fmt.Printf("Remaining Tickets: %v\n", availableTicket)

}
