package logic

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Structure to represent each card containing question and answer
type Card struct {
	Front string
	Back  string
}

// Structure to represent the deck and overall information
type Deck struct {
	Name         string
	Items        []string
	Cards        []Card
	CurrentIndex int
	filePath     string
}

func NewDeck() *Deck {
	return &Deck{
		Cards:        []Card{},
		CurrentIndex: 0,
	}
}

type DeckManager interface {
	LoadDeckMeta(filePath string) (*Deck, error)
	LoadDeckData() error

	Index() int
	Length() int

	CurrentCard() Card

	NextCard()
	PrevCard()
}

func (deck *Deck) LoadDeckMeta(filePath string) (*Deck, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create a file CSV reader instance from the file path
	reader := csv.NewReader(file)
	reader.Comma = ';'

	// Read the first line to get the headers
	headers, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading headers:", err)
		return nil, err
	}

	// Assign metadata to the deck
	deck.Items = headers
	deck.filePath = filePath

	return deck, nil
}

func (deck *Deck) LoadDeckData() error {

	// Create a file CSV reader instance from the file path
	file, err := os.Open(deck.filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	// Read the first line to get the headers
	_, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading headers:", err)
		return err
	}

	// Read the rest of the lines to get the card data
	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file or error
		}
		card := Card{
			Front: record[0],
			Back:  record[1],
		}
		deck.Cards = append(deck.Cards, card)
	}

	return nil
}

func (deck *Deck) NextCard() {
	if deck.CurrentIndex < len(deck.Cards)-1 {
		deck.CurrentIndex++
	}
}

func (deck *Deck) PrevCard() {
	if deck.CurrentIndex > 0 {
		deck.CurrentIndex--
	}
}

func (deck *Deck) Index() int {
	return deck.CurrentIndex
}

func (deck *Deck) CurrentCard() Card {
	return deck.Cards[deck.CurrentIndex]
}

func (deck *Deck) Length() int {
	return len(deck.Cards)
}
