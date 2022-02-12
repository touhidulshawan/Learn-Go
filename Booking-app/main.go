package main

import (
	"fmt"
	"strings"
)

// greet the user
func greetUser(conferenceName string, availableTicket uint32) {
	fmt.Printf("Buy tickets for current conference: %v\n", conferenceName)
	fmt.Printf("Available Tickets : %v\n", availableTicket)
	fmt.Println("----------------------------------------")
}

// get all necessary information from user to book tickets
func getUserInformation() (userFullName string, userEmail string, ticketToBuy uint32) {

	fmt.Print("Enter your first name: ")
	var firstName string
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	var lastName string
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&userEmail)

	fmt.Print("How many tickets you want to buy: ")
	fmt.Scan(&ticketToBuy)

	userFullName = firstName + " " + lastName

	return
}

func ticketConfirmationMessage(fullName string, userEmail string, ticketToBuy uint32) {

	fmt.Println("\n<-------- Confirmation Message -------->")
	fmt.Printf("Thanks %v for buying %v tickets.\n", fullName, ticketToBuy)
	fmt.Printf("Check your mailbox [%v] for further details.\n", userEmail)
}

func getFirstNames(bookings []string) []string {

	// print only first name of bookins user
	// like [Touhidul Shawan] we only use Touhidul
	// Filed splits the string with empty whitespace here
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func main() {
	fmt.Println("Welcome to our Booking App")

	conferenceName := "Love Yourself"
	var availableTicket uint32 = 25
	bookings := []string{}

	for {
		greetUser(conferenceName, availableTicket)

		var userFullName, userEmail, ticketToBuy = getUserInformation()

		if ticketToBuy <= availableTicket {
			availableTicket = availableTicket - ticketToBuy
			bookings = append(bookings, userFullName)

		} else {
			fmt.Printf("Invalid tickets number. We have only %v tickets available\n", availableTicket)
			continue
		}

		ticketConfirmationMessage(userFullName, userEmail, ticketToBuy)

		// print all bookings
		firstNames := getFirstNames(bookings)
		fmt.Printf("All Bookings : %v\n", firstNames)

		if availableTicket == 0 {
			fmt.Printf("All Tickets Sold Out For %v Conference\n", conferenceName)
			break
		}

	}
}
