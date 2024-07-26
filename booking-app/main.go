package main

import (
	"booking-app/helper"
	"fmt"
)

const conferenceTickets int = 50
var conferenceName string = "Go Conference"
var remainingTickets uint = 50

var bookings = make([]UserData, 0) //it's struct
type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main(){
	greetUsers()

	for  remainingTickets > 0 && len(bookings) < 50{
		firstName, lastName, email, userTickets := getUserInputs()
		isValidEmail, isValidName, isValidTicketNumber := helper.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber{
			bookeTicket(firstName,lastName, email, userTickets)

			firstNames := getFirstName()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			// var noTicketsRemaining bool = remainingTickets == 0
			if  remainingTickets == 0{
				// end of the program
				fmt.Println("Our Conference is booked out. Come back in next year")
				break
			}
		}else{
			if !isValidName {
				fmt.Println("Your first name of last name is too short")
			}else if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign.")
			}else if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}
		}
	}
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have to total of %v tickets and %v are available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend.")
}

func getFirstName() []string{
	firstNames := []string{}
	for _, booking := range bookings{  // '_' indicate the unused variable in Go. We use this when we don't know the uses of some variables
		// var names = strings.Fields(booking) //
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInputs()(string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint
		
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter numbers of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookeTicket(firstName string,lastName string, email string, userTickets uint ){
	remainingTickets = remainingTickets - userTickets
	
	// create a custom struct data initial
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings are %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName,userTickets, email)
	fmt.Printf("%v tickets are remaining!\n",remainingTickets)
}