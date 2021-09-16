package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Error("Expected first card of Ace of Spades, but got", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Error("Expected first card of Four of Clubs, but got", d[len(d)-1])
	}
}


func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
    os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck  := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Error("Expected 16 cards in deck, got", len(loadedDeck))
	}

	os.Remove("_decktesting")
}