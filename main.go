package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/crudenesss/interactive-cards/internal/ui"
)

func main() {
	myApp := app.NewWithID("com.crudenesss.interactivecards")
	appIcon, err := fyne.LoadResourceFromPath("assets/icon.png")
	if err == nil {
		myApp.SetIcon(appIcon)
	} else {
		log.Println("Warning: Could not load app icon:", err)
	}

	myWindow := myApp.NewWindow("Interactive Learning Cards")
	myWindow.Resize(fyne.NewSize(500, 400))

	appUI := ui.NewAppUI(myApp, myWindow)
	_ = appUI

	myWindow.SetMaster()
	myWindow.ShowAndRun()
}
