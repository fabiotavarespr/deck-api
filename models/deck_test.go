package models

import (
	"strings"
	"testing"

	internalerrors "deck-api/errors"

	"github.com/stretchr/testify/assert"
)

func TestDeckModel(t *testing.T) {
	t.Run("Get number of remaining cards successfully", func(t *testing.T) {
		codes := []string{"AS", "7D", "0H", "8D"}
		deck := Deck{Cards: codes}

		remaining := deck.GetRemaining()

		assert.Equal(t, len(codes), remaining)
	})

	t.Run("Shuffle cards successfully", func(t *testing.T) {
		deck := Deck{Cards: []string{"AS", "7D", "0H", "8D"}}

		deck.Shuffle()

		assert.True(t, deck.IsShuffled)
	})

	t.Run("Create a standard deck (with 52 cards and not shuffled) successfully", func(t *testing.T) {
		var emptyCardsSlice []Card
		standardDeck, err := NewDeck(emptyCardsSlice)

		allPossibleCards := []string{
			"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "0S", "JS", "QS", "KS",
			"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "0D", "JD", "QD", "KD",
			"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "0C", "JC", "QC", "KC",
			"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "0H", "JH", "QH", "KH",
		}

		expectedCards := strings.Join(allPossibleCards, ",")
		actualCards := strings.Join(standardDeck.Cards, ",")

		assert.NoError(t, err)
		assert.Equal(t, 52, len(standardDeck.Cards))
		assert.False(t, standardDeck.IsShuffled)
		assert.Equal(t, expectedCards, actualCards)
	})

	t.Run("Create a custom deck successfully", func(t *testing.T) {
		cardCodes := []string{"AS", "2S", "5C", "8H", "QH"}
		var cards []Card
		for _, code := range cardCodes {
			cards = append(cards, Card{Code: code})
		}

		customDeck, err := NewDeck(cards)

		expectedCards := strings.Join(cardCodes, ",")
		actualCards := strings.Join(customDeck.Cards, ",")

		assert.NoError(t, err)
		assert.Equal(t, len(cards), len(customDeck.Cards))
		assert.Equal(t, expectedCards, actualCards)
	})

	t.Run("Create a custom deck with error because there are invalid cards", func(t *testing.T) {
		invalidCardCodes := []Card{{Code: "10S"}, {Code: "8H"}, {Code: "10C"}, {Code: "10H"}, {Code: "*H"}, {Code: "invalid"}}

		customDeck, err := NewDeck(invalidCardCodes)

		assert.Error(t, err)
		assert.IsType(t, internalerrors.ErrInvalidEntry{}, err)
		assert.Empty(t, customDeck.Cards)
	})
}
