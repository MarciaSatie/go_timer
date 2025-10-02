package main

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("Timer starting...")

	// --------- Initiating Fyne ------------------------------------------------
	myApp := app.New()                      // create a new Fyne app
	myWindow := myApp.NewWindow("Go Timer") // create a window
	myWindow.Resize(fyne.NewSize(320, 180)) // ensure visible size

	// Entry where user types seconds
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter seconds")

	// Label to show countdown
	label := widget.NewLabel("Waiting for input...")

	// Button to start timer
	button := widget.NewButton("Start Timer", func() {
		// Convert entry text (string) to int
		// Go lets you catch both return values at once by listing two variables separated by commas:
		//   - secs will hold the integer (result)
		//   - err will hold the error (nil if OK)
		//
		// strconv.Atoi() comes from Go’s strconv package (“string conversion”).
		// Atoi means Ascii TO Integer.
		// It converts a string like "5" → 5 (int).
		// If the string is not a valid number (like "hello"), it returns an error.
		secs, err := strconv.Atoi(entry.Text)

		if err != nil {
			label.SetText("❌ Please enter a number")
			return
		}

		// Start a goroutine so UI doesn’t freeze
		go func(seconds int) {
			for i := seconds; i > 0; i-- {
				time.Sleep(1 * time.Second) // pause 1 second between updates
				label.SetText(fmt.Sprintf("Seconds left: %d", i-1))
			}
			label.SetText("⏰ Time's up!")
		}(secs)
	})

	// Put entry, button, and label in a vertical layout
	myWindow.SetContent(container.NewVBox(entry, button, label))
	myWindow.ShowAndRun()
}
