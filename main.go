package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("Timer starting...")

	seconds := 5

	myApp := app.New()                      // create a new Fyne app
	myWindow := myApp.NewWindow("Go Timer") // create a window
	myWindow.Resize(fyne.NewSize(320, 180)) // ensure visible size

	label := widget.NewLabel(fmt.Sprintf("Seconds left: %d", seconds))
	myWindow.SetContent(label)

	ticker := time.NewTicker(1 * time.Second)

	// Run countdown in background so UI stays responsive
	go tictac(seconds, label, ticker)

	myWindow.ShowAndRun()
}

// tictac now receives everything it needs
func tictac(seconds int, label *widget.Label, ticker *time.Ticker) {
	defer ticker.Stop() // clean up when done, defer will run before the function ends

	for i := seconds; i > 0; i-- {
		<-ticker.C
		label.SetText(fmt.Sprintf("Seconds left: %d", i-1))
	}
	label.SetText("⏰ Time's up!")
}
