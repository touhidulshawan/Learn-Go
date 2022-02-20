package main

import (
	"booking-app/validation"
	"fmt"
	"time"
    "sync"
)

var conferenceName = "Love Yourself"
var availableTicket uint32 = 25
var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint32
}

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

// function to book ticket
func bookingTicket(firstName string, lastName string, userEmail string, ticketToBuy uint32) {
	availableTicket = availableTicket - ticketToBuy

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           userEmail,
		numberOfTickets: ticketToBuy,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)
}

// get all firstname of those who booked tickets
func getFirstNames() []string {

	// print only first name of bookins user
	// like [Touhidul Shawan] we only use Touhidul
	// Filed splits the string with empty whitespace here
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

// function to send ticket after 10 second

func sendTicket(firstName string, lastName string, email string, ticketsToBuy uint32) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", ticketsToBuy, firstName, lastName)
	fmt.Println("----------------------------------")
	fmt.Printf("Tickets sent : %v to user : %v\n", ticket, email)
	fmt.Println("----------------------------------")
    wg.Done()
}

func main() {

	for {
		greetUser()

		firstName, lastName, userEmail, ticketToBuy := getUserInformation()
		isValidName, isValidEmail, isValidTicketNumber :=
			validation.UserInputValidation(firstName, lastName, userEmail, ticketToBuy, availableTicket)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookingTicket(firstName, lastName, userEmail, ticketToBuy)

            wg.Add(1)
			go sendTicket(firstName, lastName, userEmail, ticketToBuy)

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
        wg.Wait()
}
