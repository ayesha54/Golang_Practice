package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("expected deck length of 16 , but got %v", len(d))
	}

	if d[0] != "Ace of spades" {
		t.Errorf("Expected ace of spades but got %v", d[0])
	}

	if d[len(d)-1] != "five of clubs" {
		t.Errorf("Expected Four of Clubs but got %v", d[len(d)-1])
	}
}
