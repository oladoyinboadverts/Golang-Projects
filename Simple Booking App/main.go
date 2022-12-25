package main

import (
	"fmt"
)

func main() {

	var firstName string
	var lastName string
	var Ticket int
	var Email string
	//available tickets
	totalTickets := 500

	//welcome message
	fmt.Println("Welcome To My Booking App")
	fmt.Println("Please Enter Your Firstname, Last Name And The Amount Of Tickets You Want:")

	//for to repeat process again over and over till the ticket finish
	for {

		//enter username
		fmt.Println("Enter Firstname")
		fmt.Scanln(&firstName)

		//second input
		fmt.Println("Enter LastName")
		fmt.Scanln(&lastName)

		//email input
		fmt.Println("Enter Your Email:")
		fmt.Scanln(&Email)

		//tickets input
		fmt.Println("Enter The Amount Of Ticket You Are Buying:")
		fmt.Scanln(&Ticket)

		if totalTickets < Ticket || totalTickets == 0 {
			fmt.Println("Ticket is not Available Or We don't have that much Tickets Right Now")
			return
		}

		//calculate the remaining ticket after the user input how much he is buying
		remainingTickets := totalTickets - int(Ticket)

		//set RemainingTickets
		totalTickets = remainingTickets

		fmt.Printf("Your Booking Has Been Processed, Your Reservation Name Is %v %v, While The Remaining Information Will Be Sent To Your Email: %v \n\n", firstName, lastName, Email)
	}
}
