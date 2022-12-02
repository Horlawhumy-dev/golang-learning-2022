package main

import (
	"booking_app/helper/helper"
	"fmt"
	"sync"
)

// Package level params
var confTickets uint = 50
var remTickets uint = confTickets
var conferenceName = "Go Conference"
var bookingsInfo = make([]UserData, 0) // make([]map[string]string, 0) making slice of maps
var globalVarFromHelper string = helper.MyGlobalvar

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOfTicket uint
}

var wg = sync.WaitGroup{}

func main() {

	//greet
	greetUser(conferenceName)

	fmt.Println("Kindly provide the following details.")

	userFirstName, userLastName, userEmailAddress, userTicketNum := getUserDetails()

	if userTicketNum > remTickets {
		fmt.Printf("We only have %d tickets remaining. So you can't book %d tickets\n", remTickets, userTicketNum)
		// continue
	}

	isValidInputs := helper.IsValidEmailAddress(userEmailAddress) && helper.IsValidNames(userFirstName, userLastName) && helper.IsValidTicketNumber(userTicketNum, remTickets)

	if isValidInputs {
		remTickets = makeBooking(userTicketNum,
			userLastName, userFirstName, userEmailAddress)
	} else {
		if !helper.IsValidEmailAddress(userEmailAddress) {
			fmt.Print("Email address is not valid.\n")
		}

		if !helper.IsValidNames(userFirstName, userLastName) {
			fmt.Print("Invalid. First or Last name is less than two characters.")
		}

		if !helper.IsValidTicketNumber(userTicketNum, remTickets) {
			fmt.Print("Tickets number cannot be greater than the remaining tickets.")
		}
	}

	if remTickets == 0 {
		fmt.Print("Tickets finished for the conf.\n")
		// break
	}

	wg.Wait() // blocks the thread until waitgroup finishes executing goroutines
}

func greetUser(conferenceName string) {
	fmt.Printf("Welcome to %s booking application.\n", conferenceName)
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

	var userData = UserData{
		firstName:      userFirstName,
		lastName:       userLastName,
		email:          userEmailAddress,
		numberOfTicket: userTicketNum,
	}

	// userData["numberOfTicket"] = strconv.FormatUint(uint64(userTicketNum), 10) convert uint to string to uint
	//fmt.Sprint(userTicketNum)

	bookingsInfo = append(bookingsInfo, userData)
	wg.Add(1) // add the number of goroutines

	go helper.SendTicket(userEmailAddress, userFirstName, userLastName, userTicketNum) // execute in a separate thread now

	fmt.Printf("List of data: %v\n", bookingsInfo)

	firstNames := extractFirstNames

	fmt.Printf("%s are names of all bookings for the %s.\n", firstNames, conferenceName)

	fmt.Printf("%d tickets remaining for the %s.\n", remTickets, conferenceName)

	return remTickets
}

func extractFirstNames() []string {
	firstNames := []string{} // slice of string

	for _, bookingInfo := range bookingsInfo {
		//strings.Fields(bookingInfo) string splitting
		firstNames = append(firstNames, bookingInfo.firstName)
	}
	return firstNames
}
