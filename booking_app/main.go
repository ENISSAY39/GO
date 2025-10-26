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

	greetUsers()

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var UserTickets int

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

		// Validate user input

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := UserTickets > 0 && UserTickets <= int(remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - uint(UserTickets)

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation e-mail at %v															", firstName, lastName, UserTickets, email)
			fmt.Printf(" %v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

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
	fmt.Println("Welcome to our application")
}
