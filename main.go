package main

import (
	"booking-app/sharedlib"
	"fmt"
	"strings"
	//"slices"
)

var ConfData = make([]ConferenceEvent, 0)
var BookingData = make([]ConfUser, 0)

type ConferenceEvent struct {
	conferenceName       string
	conferenceLocation   string
	conferenceTickets    uint
	remainingconfTickets uint
	Confeventuser        []ConfUser
}
//Methods of the Conference event
func (CE ConferenceEvent) makebooking(userFirstName string, userLastName string, userEmail string, usertickets uint, remainingconftickets uint) ConfUser {
	remainingconftickets -= usertickets
	//storing booking data to a customer profile
	var userData = ConfUser{
		firstName:   userFirstName,
		lastName:    userLastName,
		userEmail:   userEmail,
		usertickets: usertickets,
	}
	fmt.Printf("Thanks %v %v for booking %v ticket(s). The ticket will be emailed on %v \n", userFirstName, userLastName, usertickets, userEmail)
	fmt.Printf("%v tickets remaining\n", remainingconftickets)
	fmt.Println("***************************************************")
	return userData
}

type ConfUser struct {
	firstName   string
	lastName    string
	userEmail   string
	usertickets uint
}

func main() {
	var userChoice string
	for {
		// greet the user and ask what the user wants to do
		fmt.Println("Welcome to the conference app ")
		fmt.Println("What do you want to do ?")
		fmt.Println("Press 1 to create a new conference ")
		fmt.Println("Press 2 to book a ticket in an available conference")
		fmt.Println("Press 3 to view booking details")
		fmt.Println("Print 4 to view all conference details")
		fmt.Scan(&userChoice)

		// in case statement
		switch userChoice {
		case "1":
			var confname string
			fmt.Println("1 pressed")
			fmt.Println("We will create a new conference for you, Enter the conference Name")
			fmt.Scan(&confname)

			if ValidateConfName(confname) {
				CreateConference(confname)
			} else {
				fmt.Println("Conference with the same name exists, please create a conference with a different name")
				fmt.Print("***************************************************")
			}

		case "2":
			var confnametobook string
			fmt.Println("Following conferences are available")
			printConfNames()
			fmt.Println("Enter the name of the conference you want to book ticket for")
			fmt.Scan(&confnametobook)

			for k, conf := range ConfData {
				if confnametobook == conf.conferenceName && conf.remainingconfTickets > 0 {
					sharedlib.GreetUser(conf.conferenceName, conf.conferenceTickets, conf.remainingconfTickets)
					userFirstName, userLastName, userEmail, userTicket := getuserdata()
					var isuserdatavalid, issue = validateUserData(userFirstName, userLastName, userEmail)
					if !isuserdatavalid {
						fmt.Printf("Please submit valid data, issue is in: %v \n", issue)
						fmt.Print("***************************************************")

					} else if userTicket > conf.remainingconfTickets {
						fmt.Printf("Tickets unavailable for the required amount \n")
					} else {
						var pointertoconfuser = ConfUser{}
						pointertoconfuser = conf.makebooking(userFirstName, userLastName, userEmail, userTicket, conf.remainingconfTickets)
						conf.Confeventuser = append(conf.Confeventuser, pointertoconfuser)
						ConfData[k] = conf
						fmt.Printf("The name of conference booked is %v \n", conf.conferenceName)
						fmt.Printf("Number of tickets remaining are %v \n", conf.remainingconfTickets)
						fmt.Printf("Booking details till now are %v \n", conf.Confeventuser)
						fmt.Printf("confData: %v\n", ConfData)
					}
				}
			}

		case "3":
			var confnametoview string
			fmt.Println("Following conferences are available")
			printConfNames()
			fmt.Println("Enter the name of the conference you want to view tickets for")
			fmt.Scan(&confnametoview)

			for _, conf := range ConfData {
				if confnametoview == conf.conferenceName {
					for _, user := range conf.Confeventuser {
						fmt.Printf("The user %v has booked %v tickets \n", user.firstName+""+user.lastName, user.usertickets)
					}
				}
			}
		case "4":
			printConfNames()

		default:
			fmt.Println("Press a valid option")

		}

	}
}

/*
func makeconferencebooking(conf ConferenceEvent, k int){
	var pointertoconfuser = ConfUser{
	}
	conf.remainingconfTickets,pointertoconfuser = makeBooking(userFirstName, userLastName, userEmail, userTicket,conf.remainingconfTickets)
	conf.Confeventuser = append(conf.Confeventuser, pointertoconfuser)
	ConfData[k] =conf
	fmt.Printf("The name of conference booked is %v \n",conf.conferenceName)
	fmt.Printf("Number of tickets remaining are %v \n",conf.remainingconfTickets)
	fmt.Printf("Booking details till now are %v \n",conf.Confeventuser)
	fmt.Printf("confData: %v\n", ConfData)
}
*/

func printConfNames() {
	for i, conf := range ConfData {
		fmt.Printf("Name of the %v conference is %v\n", i+1, conf.conferenceName)
	}
}
func ValidateConfName(confName string) bool {

	var isconfnameValid = true
	for _, conf := range ConfData {
		if confName == conf.conferenceName {
			isconfnameValid = false
			return isconfnameValid
		}

	}
	return isconfnameValid
}

func CreateConference(confname string) {
	var confdat = ConferenceEvent{
		conferenceName: confname,
	}
	fmt.Println("Enter the location of the conference")
	fmt.Scan(&confdat.conferenceLocation)
	fmt.Println("Enter the number of ticekts in the conference")
	fmt.Scan(&confdat.conferenceTickets)
	confdat.remainingconfTickets = confdat.conferenceTickets
	ConfData = append(ConfData, confdat)
}

func makeBooking(userFirstName string, userLastName string, userEmail string, usertickets uint, remainingTickets uint) (uint, ConfUser) {

	remainingTickets -= usertickets
	//storing booking data to a customer profile
	var userData = ConfUser{
		firstName:   userFirstName,
		lastName:    userLastName,
		userEmail:   userEmail,
		usertickets: usertickets,
	}
	fmt.Printf("Thanks %v %v for booking %v ticket(s). The ticket will be emailed on %v \n", userFirstName, userLastName, usertickets, userEmail)
	fmt.Printf("%v tickets remaining\n", remainingTickets)
	fmt.Println("***************************************************")
	return remainingTickets, userData
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

/*
func secondmain() {
}
*/
