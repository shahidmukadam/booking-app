package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go conference"

const conferenceTickets = 50

var remainingTickets uint = 50

func main() {

	var userFirstName string
	var userLastName string
	var userEmail string
	var usertickets uint
	for {
		if remainingTickets > 0 {
			greetUser()

			userFirstName, userLastName, userEmail, usertickets = getuserdata()
			var isuserdatavalid, issue = validateUserData(userFirstName, userLastName, userEmail)

			//fmt.Printf("The user data provided is %v and issue observed is in %v \n", isuserdatavalid , issue)
			if !isuserdatavalid {
				fmt.Printf("Please submit valid data, issue is in: %v \n", issue)
				fmt.Print("***************************************************")

			} else if usertickets > remainingTickets {
				fmt.Printf("Tickets unavailable for the required amount \n")
			} else {
				makeBooking(userFirstName, userLastName, userEmail, usertickets)
			}

		} else {
			break
		}

	}

}

func makeBooking(userFirstName string, userLastName string, userEmail string, usertickets uint) {
	remainingTickets -= usertickets
	fmt.Printf("Thanks %v %v for booking %v ticket(s). The ticket will be emailed on %v \n", userFirstName, userLastName, usertickets, userEmail)
	fmt.Printf("%v tickets remaining\n", remainingTickets)
	fmt.Print("***************************************************")
}
func getuserdata() (string, string, string, uint) {

	var userFirstName string
	var userLastName string
	var userEmail string
	var usertickets uint

	fmt.Println("Enter user first Name")
	fmt.Scan(&userFirstName)

	fmt.Println("Enter user Last Name")
	fmt.Scan(&userLastName)

	fmt.Println("Enter user email")
	fmt.Scan(&userEmail)

	fmt.Println("Enter how many tickets you want to book?")
	fmt.Scan(&usertickets)

	return userFirstName, userLastName, userEmail, usertickets

}
func validateUserData(userFirstName string, userLastName string, userEmail string) (bool, string) {
	var isvalid = false
	var issue = ""
	if strings.Count(userFirstName, "") >= 3 {

		if strings.Count(userLastName, "") >= 3 {
			if strings.Contains(userEmail, "@") {
				isvalid = true
			} else {
				issue = "userEmail"
				return isvalid, issue

			}

		} else {
			issue = "userLastName"
			return isvalid, issue
		}
	} else {
		issue = "userFirstName"
		return isvalid, issue
	}

	return isvalid, issue
}

func greetUser() {

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v and %v tickets are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Book your tickets with us now")
}
