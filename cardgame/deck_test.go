package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 36 {
		t.Errorf("Expected deck length of 36, but got %v", len(d))
	}

	if d[0] != "Eichel6" {
		t.Errorf("Expected first card Eichel6 but got %v", d[0])
	}

	if d[len(d)-1] != "Schellenass" {
		t.Errorf("Expected first card Schellenass but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	deck := newDeck()
	os.Remove("_decktesting")
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 36 {
		t.Errorf("Expected deck length of 36, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
