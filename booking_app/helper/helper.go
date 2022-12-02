package helper

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{} // waits for collection of goroutines to finish

var MyGlobalvar string = "anything" //global package level variable

func SendTicket(emailAddress string, firstName string, lastName string, ticketNum uint) {
	time.Sleep(5 * time.Second)
	fmt.Println("#####################")
	var ticket = fmt.Sprintf("%v tickets sent to %v %v.\n", ticketNum, firstName, lastName)
	fmt.Printf("Sending Tickets: %v \n through email address %v.\n", ticket, emailAddress)
	fmt.Println("#####################")
	wg.Done() // decrements the number of goroutines set
}

func IsValidTicketNumber(ticketsNum uint, remTickets uint) bool {
	return ticketsNum > 0 && ticketsNum <= remTickets
}

func IsValidEmailAddress(emailAddress string) bool {
	return strings.Contains(emailAddress, "@")
}

func IsValidNames(firstName string, lastName string) bool {
	return len(firstName) >= 2 && len(lastName) >= 2
}
