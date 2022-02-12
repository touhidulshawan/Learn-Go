package main

import "strings"

func userInputValidation(firstName string, lastName string, userEmail string, ticketToBuy uint32) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := ticketToBuy > 0 && ticketToBuy <= availableTicket

	return isValidName, isValidEmail, isValidTicketNumber
}
