package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/crudenesss/interactive-cards/internal/logic"
)

type AppUI struct {

	// UI components
	app fyne.App
	win fyne.Window

	deckLabel    *widget.Label
	cardPanel    *fyne.Container
	deckProgress *widget.ProgressBar
	frontText    *widget.RichText
	backText     *widget.RichText
	loadButton   *widget.Button
	answerButton *widget.Button
	nextButton   *widget.Button
	prevButton   *widget.Button

	// Logic components
	currentDeck logic.DeckManager
}

// Initialize the UI instance
func NewAppUI(app fyne.App, win fyne.Window) *AppUI {

	ui := &AppUI{
		app: app,
		win: win,
	}
	ui.loadMenu()
	// ui.updateCardDisplay() // Initial state
	return ui
}

// Init UI components to the window
func (ui *AppUI) loadMenu() {

	ui.deckLabel = widget.NewLabel("No deck loaded")
	ui.deckLabel.Alignment = fyne.TextAlignCenter

	ui.loadButton = widget.NewButton("Load Deck", ui.loadFileDialog)

	content := container.NewVBox(ui.deckLabel, ui.loadButton)

	ui.win.SetContent(content)
}
