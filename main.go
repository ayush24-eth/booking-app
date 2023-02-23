package main

import (
	//"booking-app/helper"
	"fmt"
	"time"
	"sync"
)

const conferenceTicket = 100 //const aur typecasting me : nii lagta hai
var conferenceName = "Blockchain Technology"
var remainingTicket  uint = 100
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
	isOpterForInNewsLetter bool
}
var wg = sync.WaitGroup{}

func main() {
	// fmt.Println("Hello World!")
	// age := 62
	// name := "ayush"
	// fmt.Printf("my age is %v\n", age)
	// fmt.Printf("my age is %T\n", age)
	// fmt.Printf("my name is: %v\n", name)
	// fmt.Printf("my name is: %T\n", name)
	// fmt.Printf("conference ticket is %T, remainingTicket is %T, conferenceName is %T\n", conferenceTicket, remainingTicket, conferenceName) 

	// for i:= 0; i < 10; i++ {
		// 	fmt.Printf("i am %v\n", i + 1)
		// }
		
		// arr :=[]int{1,2,3,5}
		// fmt.Print(arr)
		// var bookings [50] string //array
		// bookings[0] = firstName+" "+lastName // to insert in an array
		// fmt.Printf("Length of arr is: %v\n", len(bookings))
		// fmt.Println(&remainingTicket) // address of variable
		

	greetUser()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicket{

		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)

		firName := printFirstName()
		fmt.Printf("Array of FirstNames are: %v\n", firName)

		if remainingTicket == 0 {
			fmt.Println("SOLD OUT")
		}
	} else {
		if !isValidName {
			fmt.Println("Name is too short, Enter minimun 3 characters")
		}
		if !isValidEmail {
			fmt.Println("Invalid Email")
		}
		if !isValidTicket {
			fmt.Println("No of tickets entered is invaid")
		}
		// fmt.Printf("We only have %v tickets\n", remainingTicket)
		// fmt.Printf("Your input data is invalid, try again")
	}

	// city :="London"
	// switch city {
	// 	case "NYC", "California":

	// 	case "Singapore":

	// 	case "London":

	// 	case "Delhi", "Mumbai":

	// 	case "Dubai":

	// 	default:
	// 		fmt.Println("No valid city")

	// }
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to the %v Conference\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your to tickets to attend")
}

func printFirstName() []string{
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// var firstName = names[0]
		firstNames = append(firstNames, booking.firstName)
	}
	// fmt.Printf("These are all the bookings %v\n", bookings)
	// fmt.Printf("The first names of bookings are: %v\n", firstNames)
	return firstNames
}

func getUserInput() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter no of tickets: ") 
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTicket -= userTickets
	var userData = userData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData) // to insert in slice
	fmt.Printf("List of bookings %v", bookings)

	fmt.Printf("Thankyou %v %v you have booked %v tickets. You will get a confirmation mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v Conference\n", remainingTicket, conferenceName)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10*time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:\n %v \n to email address %v \n", ticket, email)
	fmt.Println("##################")
	wg.Done()
	
}