package main

import (
	"fmt"
	"strconv"
	"strings"
)

var conferenceName = "Go Conference"
var conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	//welcome message
	greetUsers()

	//for loop to ask to book tickets. store the name+surname in a slice.
	for {

		// get user input
		firstName, lastName, email, userTickets := getUserInput()

		//check the name, email, amount of tickets
		isValidName, isValidEmail, isValidTicketNumber := inputValidation(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			//book a ticket
			bookTicket(userTickets, firstName, lastName, email)

			// call function to iterate through bookings list to print only names
			firstNames := getFirstNames()
			fmt.Printf("the first names of our bookings: %v\n", firstNames)

			//break when user tries to overbook
			if remainingTickets == 0 {
				//end program
				fmt.Println("our conference is booked up. come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your name or surname is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email is wrong")
			}
			if !isValidTicketNumber {
				fmt.Println("Your value of tickets is invalid")
			}
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("Get your ticket here to attend")
	fmt.Printf("we have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
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
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings is %v\n", bookings)

	fmt.Printf("thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func inputValidation(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
