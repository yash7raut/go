package main

import ("fmt"
		"strings"
)


func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("conferenceTickets is %T, conferenceName is %T, remainingTickets is %T \n", conferenceTickets, conferenceName, remainingTickets)
	fmt.Println()
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets, out of which %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
	fmt.Println()

	var bookings []string

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for{
		fmt.Println("Enter your first Name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last Name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email address")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets you want")
	fmt.Scan(&userTickets)
	fmt.Println()

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first element of the slice: %v\n", bookings[0])
	// fmt.Printf("The type of slice: %T\n", bookings)
	// fmt.Printf("The size of slice: %v\n", len(bookings))

	fmt.Println()
	fmt.Printf("Thankyou %v %v for booking %v tickets.You will get a confirmation mail at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	
	firstNames := []string{} 
	for index,booking := range bookings{
		var names = strings.Fields(booking)
		var firstName = names[0]
	}
	fmt.Printf("These are all our bookings: %v\n",bookings)

	}
}
