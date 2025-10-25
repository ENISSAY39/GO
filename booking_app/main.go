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

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var UserTickets int

		// Ask user first and last name

		fmt.Println("Enter your first name and last name:")
		fmt.Scan(&firstName)
		fmt.Scan(&lastName)

		// Ask user email address

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		// Ask user number of tickets

		fmt.Println("Enter your number of tickets:")
		fmt.Scan(&UserTickets)

		if UserTickets > int(remainingTickets) {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, UserTickets)
			continue
		}

		remainingTickets = remainingTickets - uint(UserTickets)

		// Store user data

		bookings = append(bookings, firstName+" "+lastName)

		// Print user data

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

	}

}
