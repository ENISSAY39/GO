package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {
	greetUsers()

	for {
		// Get user input
		firstName, lastName, email, UserTickets := getUserInput()
		// Validate user input
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, UserTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// Book ticket
			bookTicket(UserTickets, firstName, lastName, email)

			// call function to print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered doesn't contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid.")
			}

			fmt.Println("Your input data 	is invalid. Please try again.")
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var UserTickets uint

	// Ask user first and last name

	fmt.Println("Enter your first name :")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name :")
	fmt.Scan(&lastName)

	// Ask user email address

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	// Ask user number of tickets

	fmt.Println("Enter your number of tickets:")
	fmt.Scan(&UserTickets)

	return firstName, lastName, email, UserTickets
}

func bookTicket(UserTickets uint, firstName string, lastName string, email string) ([]string, uint) {
	remainingTickets = remainingTickets - uint(UserTickets)

	// create a map for a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["NumberOfTickets"] = strconv.FormatUint(uint64(UserTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation e-mail at %v", firstName, lastName, UserTickets, email)
	fmt.Printf(" %v tickets remaining for %v\n", remainingTickets, conferenceName)

	return getFirstNames(), remainingTickets

}
