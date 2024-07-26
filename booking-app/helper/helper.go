package helper

import "strings"

// for export a function from user define package, we have to initial function name with CAPITAL latter
func ValidateUserInputs(firstName string, lastName string,email string, userTickets uint, remainingTickets uint)(bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2 // ':=' is the alternative of 'var'
	isValidEmail := strings.Contains(email,"@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidEmail, isValidName, isValidTicketNumber
}