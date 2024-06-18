package explore

import "fmt"

func GreetUser(conferenceName string, conferenceTickets int ,remainingTickets uint) {

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v and %v tickets are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Book your tickets with us now")
}