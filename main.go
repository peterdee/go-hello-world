package main

import (
	"fmt"
	"time"
)

// create a custom string type
type Message string

// create a function that shows a message
func showText(text Message) {
	fmt.Println(text)
}

// create a function that shows current date and time
func showDate() {
	date := time.Now().Format("Monday, January 02, 3:4 PM")
	fmt.Println("Current date and time:", date)
}

// run the app
func main() {
	const message Message = "Hello, GO!"

	// call showText function to print the message
	showText(message)

	// call showDate function to print the current date and time
	showDate()
}
