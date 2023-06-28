package main

import (
	"os"
	"testing"
)

func TestCreateNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(deck))
	}
}

func TestSaveAndLoadFile(t *testing.T) {
	os.Remove("_testFile")
	deck := newDeck()
	deck.saveToDisk("_testFile")

	loadedDeck := newDeckFromFile("_testFile")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v !", len(deck))
	}

	os.Remove("_testFile")
}
