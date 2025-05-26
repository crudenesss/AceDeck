package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/crudenesss/interactive-cards/internal/logic"
)

// Open a file dialog to select a CSV file
func (ui *AppUI) loadFileDialog() {

	fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, ui.win)
			return
		}
		defer reader.Close()

		filePath := reader.URI().Path()
		ui.loadDeckFromPath(filePath)

	}, ui.win)

	fileDialog.Show()
}

// Load the deck from the selected file path
func (ui *AppUI) loadDeckFromPath(filePath string) {

	newDeck, err := logic.NewDeck().LoadDeckMeta(filePath)
	if err != nil {
		dialog.ShowError(err, ui.win)
		return
	}

	err = newDeck.LoadDeckData()
	if err != nil {
		dialog.ShowError(err, ui.win)
		return
	}

	deckName := fmt.Sprintf("%s -> %s", newDeck.Items[0], newDeck.Items[1])

	ui.currentDeck = newDeck
	ui.deckLabel.SetText(deckName)
	ui.deckProgress = widget.NewProgressBar()

	ui.frontText = widget.NewRichText()
	ui.backText = widget.NewRichText()

	ui.cardPanel = container.NewStack(
		ui.frontText,
		ui.backText,
	)

	cardContainer := container.NewPadded(ui.cardPanel)

	ui.answerButton = widget.NewButton("Show Answer", ui.showAnswer)
	ui.nextButton = widget.NewButton("Next", ui.nextCard)
	ui.prevButton = widget.NewButton("Prev", ui.prevCard)

	err = ui.updateCardDisplay()
	if err != nil {
		dialog.ShowError(err, ui.win)
		return
	}

	topBar := container.NewVBox(
		ui.deckLabel,
		ui.deckProgress,
	)

	navBar := container.NewHBox(
		ui.prevButton,
		ui.nextButton,
	)
	content := container.NewVBox(topBar, cardContainer, ui.answerButton, navBar)

	ui.win.SetContent(content)
}
