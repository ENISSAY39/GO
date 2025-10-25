package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	remainingTickets = remainingTickets - uint(UserTickets)

	// Store user data

	bookings = append(bookings, firstName+" "+lastName)

	// Print user data

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation e-mail at %v															", firstName, lastName, UserTickets, email)
	fmt.Printf(" %v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings: %v\n", bookings)
}
