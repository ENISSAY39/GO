package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		// Get user input
		firstName, lastName, email, UserTickets := getUserInput()
		// Validate user input
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, UserTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// Book ticket
			bookTicket(UserTickets, firstName, lastName, email, bookings, remainingTickets, conferenceName)

			// call function to print first names
			firstNames := getFirstNames(bookings)
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

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && int(userTickets) <= int(remainingTickets)
	return isValidName, isValidEmail, isValidTicketNumber
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

func bookTicket(UserTickets uint, firstName string, lastName string, email string, bookings []string, remainingTickets uint, conferenceName string) ([]string, uint) {
	remainingTickets = remainingTickets - uint(UserTickets)

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation e-mail at %v", firstName, lastName, UserTickets, email)
	fmt.Printf(" %v tickets remaining for %v\n", remainingTickets, conferenceName)
	return bookings, remainingTickets
}
