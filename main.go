package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// Main window.
var window fyne.Window

func main() {
	// Application.
	application := app.New()
	// Create the window.
	window = application.NewWindow("Desechos biodegrabables")
	// Do the window fullscreen.
	window.SetFullScreen(true)
	// Set the window's content.
	window.SetContent(getMainActivity())
	// Show the application
	window.ShowAndRun()
}
