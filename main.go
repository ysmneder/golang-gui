package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Sample Login Page")
	myWindow.Resize(fyne.NewSize(500, 400))
	myWindow.CenterOnScreen()

	entrySize := fyne.NewSize(300, 35)

	// Create Name Input
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Name")
	nameEntry.Resize(entrySize)          // set the size of the entry
	nameEntry.Move(fyne.NewPos(100, 50)) // set the position of the entry

	// Create Password Input
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")
	password.Resize(entrySize)           // set the size of the entry
	password.Move(fyne.NewPos(100, 100)) // set the position of the entry

	nameContainer := container.NewWithoutLayout(nameEntry)

	passwordContainer := container.NewWithoutLayout(password)

	// Create a submit button and handle the form data
	submitButton := widget.NewButton("Submit", func() {
		// handle the form data after the submit button is clicked
		fmt.Println("Name: ", nameEntry.Text)
		fmt.Println("Password: ", password.Text)

	})
	submitButton.Resize(fyne.NewSize(150, 35)) // set the size of the button
	submitButton.Move(fyne.NewPos(170, 160))   // set the position of the button

	// Add the form and submit button to a container and set the container as the window's content
	myContainer := container.NewWithoutLayout(nameContainer, passwordContainer, submitButton)

	myWindow.SetContent(myContainer)
	myWindow.ShowAndRun()
}
