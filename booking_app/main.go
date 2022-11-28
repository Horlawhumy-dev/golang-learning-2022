package main

import (
	"fmt"
	"strings"
)

// Package level params
var confTickets uint = 50
var remTickets uint = confTickets
var conferenceName = "Go Conference"
var bookingsInfo = []string{}

func main() {

	//greet
	greetUser()

	for {

		fmt.Println("Kindly provide the following details.")

		userFirstName, userLastName, userEmailAddress, userTicketNum := getUserDetails()

		if userTicketNum > remTickets {
			fmt.Printf("We only have %d tickets remaining. So you can't book %d tickets\n", remTickets, userTicketNum)
			continue
		}

		isValidInputs := isValidEmailAddress(userEmailAddress) && isValidNames(userFirstName, userLastName) && isValidTicketNumber(userTicketNum)

		if isValidInputs {
			remTickets = makeBooking(userTicketNum,
				userLastName, userFirstName, userEmailAddress)
		} else {
			if isValidEmailAddress(userEmailAddress) {
				fmt.Print("Email address is not valid.\n")
			}

			if !isValidNames(userFirstName, userLastName) {
				fmt.Print("Invalid. First or Last name is less than two characters.")
			}

			if !isValidTicketNumber(userTicketNum) {
				fmt.Print("Tickets number cannot be greater than the remaining tickets.")
			}
		}

		if remTickets == 0 {
			fmt.Print("Tickets finished for the conf.\n")
			break
		}

	}

}

func greetUser() {
	fmt.Printf("Welcome to %s booking application.\n", conferenceName)
	fmt.Printf("We have a total of  %d tickets available and %d tickets are remaining for booking.\n", confTickets, remTickets)
	fmt.Println("Book your ticket here to attend.")
}

func getUserDetails() (string, string, string, uint) {
	var userFirstName string
	var userLastName string
	var userEmailAddress string
	var userTicketNum uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&userFirstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&userLastName)

	fmt.Print("What is your email address: ")
	fmt.Scan(&userEmailAddress)

	fmt.Print("How many tickets would you like to book?: ")
	fmt.Scan(&userTicketNum)

	return userFirstName, userLastName, userEmailAddress, userTicketNum
}

func makeBooking(userTicketNum uint,
	userLastName string, userFirstName string, userEmailAddress string) uint {

	fmt.Printf("Hi %s %s! You have booked for %d tickets. You will receive a confirmation message at your address %s.\n",
		userFirstName, userLastName, userTicketNum, userEmailAddress)

	confTickets = confTickets - userTicketNum
	remTickets = confTickets
	bookingsInfo = append(bookingsInfo, userFirstName+" "+userLastName)

	firstNames := extractFirstNames()

	fmt.Printf("%v are names of all bookings for the %s.\n", firstNames, conferenceName)

	fmt.Printf("%d tickets remaining for the %s.\n", remTickets, conferenceName)

	return remTickets
}

func extractFirstNames() []string {
	firstNames := []string{}

	for _, bookingInfo := range bookingsInfo {
		var userName = strings.Fields(bookingInfo)
		var firstName = userName[0]
		firstNames = append(firstNames, firstName)
	}

	return firstNames
}

func isValidTicketNumber(ticketsNum uint) bool {
	return ticketsNum > 0 && ticketsNum <= remTickets
}

func isValidEmailAddress(emailAddress string) bool {
	return strings.Contains(emailAddress, "@")
}

func isValidNames(firstName string, lastName string) bool {
	return len(firstName) >= 2 && len(lastName) >= 2
}
