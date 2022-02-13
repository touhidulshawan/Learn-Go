package main

import (
	"booking-app/validation"
	"fmt"
	"strings"
)

var conferenceName = "Love Yourself"
var availableTicket uint32 = 25
var bookings = []string{}

// greet the user
func greetUser() {
	fmt.Println("Welcome to our Booking App")
	fmt.Printf("Buy tickets for current conference: %v\n", conferenceName)
	fmt.Printf("Available Tickets : %v\n", availableTicket)
	fmt.Println("----------------------------------------")
}

// get all necessary information from user to book tickets
func getUserInformation() (firstName string, lastName string, userEmail string, ticketToBuy uint32) {

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&userEmail)

	fmt.Print("How many tickets you want to buy: ")
	fmt.Scan(&ticketToBuy)

	return
}

func ticketConfirmationMessage(firstName string, lastName string, userEmail string, ticketToBuy uint32) {

	fmt.Println("\n<-------- Confirmation Message -------->")
	fmt.Printf("Thanks %v %v for buying %v tickets.\n", firstName, lastName, ticketToBuy)
	fmt.Printf("Check your mailbox [%v] for further details.\n", userEmail)
}

func getFirstNames() []string {

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

	for {
		greetUser()

		firstName, lastName, userEmail, ticketToBuy := getUserInformation()
		isValidName, isValidEmail, isValidTicketNumber :=
			validation.UserInputValidation(firstName, lastName, userEmail, ticketToBuy, availableTicket)

		if isValidName && isValidEmail && isValidTicketNumber {
			availableTicket = availableTicket - ticketToBuy
			bookings = append(bookings, firstName+" "+lastName)
			ticketConfirmationMessage(firstName, lastName, userEmail, ticketToBuy)

			// print all bookings
			firstNames := getFirstNames()
			fmt.Printf("All Bookings : %v\n", firstNames)
		} else {
			if !isValidName {
				fmt.Println("Error: First Name or Last Name is too short!")
			} else if !isValidEmail {
				fmt.Println("Error: Invalid Email! Does not contain @ symbol")
			} else if !isValidTicketNumber {
				fmt.Printf("Error: Invalid number of Tickets to buy. Available Tickets: %v\n", availableTicket)
			}
		}

		if availableTicket == 0 {
			fmt.Printf("All Tickets Sold Out For %v Conference\n", conferenceName)
			break
		}

	}
}
