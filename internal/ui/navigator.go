package ui

// Update the display of the current card in the UI
func (ui *AppUI) updateCardDisplay() error {
	ui.prevButton.Show()
	ui.nextButton.Show()
	if ui.currentDeck.Index() == 0 {
		ui.prevButton.Hide()
	} else if ui.currentDeck.Index() == ui.currentDeck.Length()-1 {
		ui.nextButton.Hide()
	}

	ui.frontText.ParseMarkdown(ui.currentDeck.CurrentCard().Front)
	ui.backText.ParseMarkdown(ui.currentDeck.CurrentCard().Back)
	ui.frontText.Show()
	ui.backText.Hide()
	return nil
}

// Update progress bar based on the current deck state
func (ui *AppUI) updateDeckProgress() error {
	if ui.currentDeck == nil {
		return nil
	}
	progress := float64(ui.currentDeck.Index()+1) / float64(ui.currentDeck.Length())
	ui.deckProgress.SetValue(progress)
	return nil
}

// Show the answer for the current card
func (ui *AppUI) showAnswer() {
	ui.backText.Show()
	ui.frontText.Hide()
}

// Move to the next card in the deck
func (ui *AppUI) nextCard() {
	ui.currentDeck.NextCard()
	ui.updateCardDisplay()
	ui.updateDeckProgress()
}

// Move to the previous card in the deck
func (ui *AppUI) prevCard() {
	ui.currentDeck.PrevCard()
	ui.updateCardDisplay()
	ui.updateDeckProgress()
}
