package main

import (
	"booking-app/sharedlib"
	"fmt"
	"strings"
)

var conferenceName = "Go conference"

const conferenceTickets = 50

var remainingTickets uint = 50
//var confData = make([]ConferenceEvent, 0)
var bookingData =make([]ConfUser ,0)

/*
type ConferenceEvent struct {
	conferenceName srring
	conferenceLocation string
	conferenceTickets uint
	remainingconfTickets uint
}*/
type ConfUser struct{
	firstName string
	lastName string
	userEmail string
	usertickets uint
}

func main() {

	var userFirstName string
	var userLastName string
	var userEmail string
	var usertickets uint
	for {
		if remainingTickets > 0 {
			sharedlib.GreetUser(conferenceName, conferenceTickets , remainingTickets)

			userFirstName, userLastName, userEmail, usertickets = getuserdata()
			var isuserdatavalid, issue = validateUserData(userFirstName, userLastName, userEmail)

			//fmt.Printf("The user data provided is %v and issue observed is in %v \n", isuserdatavalid , issue)
			if !isuserdatavalid {
				fmt.Printf("Please submit valid data, issue is in: %v \n", issue)
				fmt.Print("***************************************************")

			} else if usertickets > remainingTickets {
				fmt.Printf("Tickets unavailable for the required amount \n")
			} else {
				remainingTickets = makeBooking(userFirstName, userLastName, userEmail, usertickets,remainingTickets)
			}

		} else {
			publishBookingdata()
			break
		}

	}

}

func makeBooking(userFirstName string, userLastName string, userEmail string, usertickets uint,remainingTickets uint) uint  {
	remainingTickets -= usertickets
	//storing booking data to a customer profile
	//var bookingData = []string{}//slice
	
	var userData = ConfUser{
		firstName: userFirstName,
		lastName: userLastName,
		userEmail: userEmail,
		usertickets: usertickets,
	}

	bookingData = append(bookingData, userData)

	fmt.Printf("Thanks %v %v for booking %v ticket(s). The ticket will be emailed on %v \n", userFirstName, userLastName, usertickets, userEmail)
	fmt.Printf("%v tickets remaining\n", remainingTickets)
	fmt.Println("***************************************************")
	fmt.Printf("List of bookings till now %v \n",bookingData)
	return remainingTickets
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
func publishBookingdata(){

	for _,booking := range bookingData{

		fmt.Printf("User %v booked %v tickets. \n",booking.firstName+" "+booking.lastName,booking.usertickets)
		fmt.Printf("Ticket details will be emailed on %v \n",booking.userEmail)

	}

}
/*
func secondmain() {

	var userChoice string
	for {
		// greet the user and ask what the user wants to do 
		fmt.Println("Welcome to the conference app ")
		fmt.Println("What do you want to do ?")
		fmt.Println("Press 1 to create a new conference ")
		fmt.Println("Press 2 to book a ticket in an available conference")
		fmt.Println("Press 3 to view booking details")
		fmt.Scan(&userChoice)

		// in case statement
		switch userChoice {
		case "1":
			fmt.Println("1 pressed")
		case "2":
			fmt.Println("2 pressed")
		case "3":
			fmt.Println("3 pressed")
		default : 
		fmt.Println("Press a valid option")
			
		}
		// Create a conference  - > Take input of the conference - > Validate and create
		//book ticket in a conference -> show available conference - > book valid tickets and store booking data
		// view conference booking details  - number of tickets booked and users who have booked tickets

	}

}
*/