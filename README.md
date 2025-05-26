# Go Interactive Cards GUI

A simple and interactive GUI application for managing and displaying cards, built with Go. Decks are generated via **_.txt/.csv_** file import.

## Installation

Download binary file from Release page.

### OR - Compile from sources

Requirements: **Go version 1.24.3** or higher.

```bash
git clone https://github.com/crudenesss/AceDeck.git
cd AceDeck
go mod tidy
go build -o <output-file>
```

## Usage

In order for AceDeck to parse your cards, prepare text file with CSV-like format with `;` as **necessary delimiter** (view the sample file by [this link](./internal/data/sample.csv))

## Roadmap

### Interface

- [ ] Implement some decent design
- [ ] Tab for managing recent opened decks

### Logic

- [ ] Cards shuffling
- [ ] Swap guessing items for recently opened decks


