package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}

var conferenceName = "Go Conference"

const conferenceTicket uint = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

func main() {
	greetUsers()
	firstName, lastName, email, tickets := getUserInput()
	isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, tickets, email)

	if isValidName && isValidEmail && isValidTickets {
		bookTicket(firstName, lastName, tickets, email)
		waitGroup.Add(1)
		go sendTicket(firstName, lastName, tickets, email)
		firstNames := getFirstNames()
		fmt.Printf("The firstnames of bookings are %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Sorry, No More Tickets Left")
		}

	} else {
		if !isValidName {
			fmt.Println("Invalid firstname or lastname given")
		}
		if !isValidEmail {
			fmt.Println("Invalid Email given")
		}
		if !isValidTickets {
			fmt.Println("Invalid ticket number is given or the given number of tickets exceeds the remaining tickets")
		}
	}
	waitGroup.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Remaining Available Tickets are %v\n", remainingTickets)
	fmt.Printf("Book your Tickets now!\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
func validateUserInput(firstName string, lastName string, tickets uint, email string) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := tickets > 0 && tickets < remainingTickets
	return isValidName, isValidEmail, isValidTickets
}
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var tickets uint

	fmt.Println("Enter first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter Last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter Email")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets")
	fmt.Scan(&tickets)

	return firstName, lastName, email, tickets
}
func bookTicket(firstName string, lastName string, tickets uint, email string) {
	remainingTickets = remainingTickets - tickets
	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   tickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v\n", bookings)

	fmt.Printf("User %v %v Booked %v Tickets you will recieve confirmation email at %v\n", firstName, lastName, tickets, email)
	fmt.Printf("Remaining tickets are %v\n", remainingTickets)
}
func sendTicket(firstName string, lastName string, tickets uint, email string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v%v", tickets, firstName, lastName)
	fmt.Println("###########")
	fmt.Printf("Sending tickets:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###########")
	waitGroup.Done()
}
