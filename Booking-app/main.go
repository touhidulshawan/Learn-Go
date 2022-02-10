package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Booking App")

	conferenceName := "Love Yourself"
	var availableTicket uint32 = 25
	bookings := []string{}
	for {

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

		if ticketToBuy <= availableTicket {
			availableTicket = availableTicket - ticketToBuy
			bookings = append(bookings, firstName+" "+lastName)

		} else {
			fmt.Printf("Invalid tickets number. We have only %v tickets available\n", availableTicket)
			continue
		}

		fmt.Println()
		fmt.Println("<-------- Confirmation Message -------->")

		fmt.Printf("Thanks %v %v for buying %v tickets.\n", firstName, lastName, ticketToBuy)
		fmt.Printf("Check your mailbox [%v] for further details.\n", userEmail)

		// print only first name of bookins user
		// like [Touhidul Shawan] we only use Touhidul
		// Filed splits the string with empty whitespace here
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("Bookings : %v\n", firstNames)
		fmt.Println("-------------------------------------------------")

		if availableTicket == 0 {
			fmt.Printf("All Tickets Sold Out For %v Conference\n", conferenceName)
			break
		}

	}
}
